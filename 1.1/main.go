package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	b, err := os.ReadFile("./1/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculate(b))
}

func calculate(b []byte) uint64 {
	var sum uint64
	var s = string(b)
	for _, line := range strings.Split(s, "\n") {
		log.Print("line: ", line)

		bytes := []byte(line)
		left := firstNumber(bytes)

		slices.Reverse(bytes)
		right := firstNumber(bytes)

		number := left*10 + right

		log.Print("number: ", number, "\n")
		sum += number
	}

	return sum
}

func firstNumber(b []byte) uint64 {
	for _, bb := range b {
		if !isNumber(bb) {
			continue
		}
		return toNumber(bb)
	}
	return 0
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func toNumber(b byte) uint64 {
	return uint64(b - '0')
}
