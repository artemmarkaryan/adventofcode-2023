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
func (s set) setred(i int32) set   { s[0] = i; return s }
func (s set) setgreen(i int32) set { s[1] = i; return s }
func (s set) setblue(i int32) set  { s[2] = i; return s }
func (s set) power() int32         { return s.blue() * s.green() * s.red() }

func parseSet(line string) (s set) {
	for _, part := range strings.Split(line, ", ") {
		split := strings.Split(strings.TrimSpace(part), " ")
		color := split[1]
		color = strings.TrimSpace(color)
		number, _ := strconv.Atoi(split[0])
		f := map[string]func(int32) set{
			"red":   s.setred,
			"green": s.setgreen,
			"blue":  s.setblue,
		}[color]
		s = f(int32(number))
	}
	return
}

func processLine(line string) (power int32) {
	split := strings.SplitN(line, ":", 2)
	righthand := split[1]

	set := set{}
	for _, rawSet := range strings.Split(righthand, ";") {
		parsed := parseSet(rawSet)

		if parsed.blue() > set.blue() {
			set = set.setblue(parsed.blue())
		}

		if parsed.green() > set.green() {
			set = set.setgreen(parsed.green())
		}

		if parsed.red() > set.red() {
			set = set.setred(parsed.red())
		}
	}

	log.Print(set)

	return set.power()
}

func main() {
	file, err := os.OpenFile("./3/input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := atomic.Int32{}
	wg := sync.WaitGroup{}
	for scanner.Scan() {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			pwr := processLine(line)
			sum.Add(pwr)
		}(scanner.Text())
	}
	wg.Wait()
	fmt.Print(sum.Load())
}
