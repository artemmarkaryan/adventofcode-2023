package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type set [3]int32

func (s set) red() int32           { return s[0] }
func (s set) green() int32         { return s[1] }
func (s set) blue() int32          { return s[2] }
func (s set) setRed(i int32) set   { s[0] = i; return s }
func (s set) setGreen(i int32) set { s[1] = i; return s }
func (s set) setBlue(i int32) set  { s[2] = i; return s }

func parseSet(line string) (s set) {
	for _, part := range strings.Split(line, ", ") {
		split := strings.Split(strings.TrimSpace(part), " ")
		color := split[1]
		color = strings.TrimSpace(color)
		number, _ := strconv.Atoi(split[0])
		f := map[string]func(int32) set{
			"red":   s.setRed,
			"green": s.setGreen,
			"blue":  s.setBlue,
		}[color]
		s = f(int32(number))
	}
	return
}

func (s set) possibleFor(other set) bool {
	return true &&
		s.blue() <= other.blue() &&
		s.green() <= other.green() &&
		s.red() <= other.red()
}

func processLine(line string, config set) (gameID int, possible bool) {
	split := strings.SplitN(line, ":", 2)
	lefthand := split[0]
	lefthandSplit := strings.SplitN(lefthand, " ", 2)
	gameID, _ = strconv.Atoi(lefthandSplit[1])
	righthand := split[1]
	for _, rawSet := range strings.Split(righthand, ";") {
		if !parseSet(rawSet).possibleFor(config) {
			return
		}
	}
	possible = true
	return
}

func main() {
	file, err := os.OpenFile("./3/input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	config := set{}.
		setRed(12).
		setGreen(13).
		setBlue(14)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := atomic.Int32{}
	wg := sync.WaitGroup{}
	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			id, ok := processLine(line, config)
			if ok {
				log.Print("id: ", id)
				sum.Add(int32(id))
			}
		}(scanner.Text())
	}
	wg.Wait()
	fmt.Print(sum.Load())
}
