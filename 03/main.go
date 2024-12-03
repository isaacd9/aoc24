package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	rex1 = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	rex2 = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don\'t\(\)`)
)

func part1(in string) int {
	matches := rex1.FindAllStringSubmatch(in, -1)
	var v int

	for _, match := range matches {
		fst, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		snd, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		log.Printf("%v", match)

		v += fst * snd
	}

	return v
}

func part2(in string) int {
	matches := rex2.FindAllStringSubmatch(in, -1)
	var (
		v  int
		do bool = true
	)

	for _, match := range matches {
		log.Printf("%v", match)
		sp := strings.Split(match[0], "(")[0]
		switch sp {
		case "mul":
			if do {
				log.Printf("enabled %v", sp)
				fst, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				snd, err := strconv.Atoi(match[2])
				if err != nil {
					panic(err)
				}

				v += fst * snd
			}
		case "do":
			do = true
		case "don't":
			do = false
		}
	}

	return v
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", part1(string(b)))
	fmt.Printf("%d\n", part2(string(b)))
}
