package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
	"sync/atomic"
)

func main() {
	file, err := os.Open("./2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := atomic.Int32{}
	wg := sync.WaitGroup{}
	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			number := calculateLine(line)
			sum.Add(number)
		}(scanner.Text())
	}
	wg.Wait()
	fmt.Println(sum.Load())
}

type token struct {
	key   string
	value int32
}

var straigntTokens = []token{
	{"zero", 0},
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
	{"0", 0},
	{"1", 1},
	{"2", 2},
	{"3", 3},
	{"4", 4},
	{"5", 5},
	{"6", 6},
	{"7", 7},
	{"8", 8},
	{"9", 9},
}

var reversedTokens = func() []token {
	var result = make([]token, len(straigntTokens))
	copy(result, straigntTokens)
	for i := range result {
		bytes := []byte(result[i].key)
		slices.Reverse(bytes)
		result[i].key = string(bytes)
	}
	return result
}()

func calculateLine(line string) int32 {
	leftmost := firstNumber(line, straigntTokens)

	bytes := []byte(line)
	slices.Reverse(bytes)
	reversedLine := string(bytes)

	rightmost := firstNumber(reversedLine, reversedTokens)

	return leftmost*10 + rightmost
}

func firstNumber(line string, tokens []token) int32 {
	var iterations = make([]int, len(tokens))

	for i := range line {
		lineByte := line[i]
		for tokenIndex, token := range tokens {
			tokenIteration := iterations[tokenIndex]
			if lineByte == token.key[tokenIteration] {
				iterations[tokenIndex] = tokenIteration + 1
			}
			if iterations[tokenIndex] == len(token.key) {
				return token.value
			}
		}
	}

	return 0
}
