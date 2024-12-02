package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func asc(row []int) bool {
	for n := 1; n < len(row); n++ {
		it := row[n]
		log.Printf("[asc] it: %v, row[n-1]: %v", it, row[n-1])
		if it-row[n-1] < 1 || it-row[n-1] > 3 {
			return false
		}
	}
	return true
}

func dsc(row []int) bool {
	for n := 1; n < len(row); n++ {
		it := row[n]
		log.Printf("[dsc] it: %v, row[n-1]: %v", it, row[n-1])
		if row[n-1]-it < 1 || row[n-1]-it > 3 {
			return false
		}
	}
	return true
}

func part1(in [][]int) int {
	var c int

	for _, row := range in {
		log.Printf("row: %v", row)
		if row[1] > row[0] {
			if asc(row) {
				c++
			}
		} else {
			if dsc(row) {
				c++
			}
		}
	}

	return c
}

func asc2(row []int) bool {
	var bad bool

	if row[1]-row[0] < 1 || row[1]-row[0] > 3 {
		log.Printf("removing %v", row[0])
		row = row[1:]
		bad = true
	}

	for n := 1; n < len(row); n++ {
		it := row[n]
		log.Printf("[asc] it: %v, row[n-1]: %v", it, row[n-1])

		if it-row[n-1] < 1 || it-row[n-1] > 3 {
			if bad {
				return false
			}
			log.Printf("removing %v, n: %+v", row[n], n)
			row = append(row[:n], row[n+1:]...)
			n--
			bad = true
		}
	}
	return true
}

func dsc2(row []int) bool {
	var bad bool

	if row[0]-row[1] < 1 || row[0]-row[1] > 3 {
		log.Printf("removing %v", row[0])
		row = row[1:]
		bad = true
	}

	for n := 1; n < len(row); n++ {
		it := row[n]
		log.Printf("[dsc] it: %v, row[n-1]: %v", it, row[n-1])
		if row[n-1]-it < 1 || row[n-1]-it > 3 {
			if bad {
				return false
			}
			log.Printf("removing %v, n: %+v", row[n], n)
			row = append(row[:n], row[n+1:]...)
			n--
			bad = true
		}
	}
	return true
}

func part2(in [][]int) int {
	var c int

	for _, row := range in {
		log.Printf("row: %v", row)
		if row[1] > row[0] {
			if asc2(row) {
				fmt.Println("ok")
				c++
			} else {
				fmt.Println("not ok")
			}
		} else {
			if dsc2(row) {
				fmt.Println("ok")
				c++
			} else {
				fmt.Println("not ok")
			}
		}
	}

	return c
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var in [][]int

	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}
		ln := strings.Split(line, " ")
		var lnn []int
		for _, it := range ln {
			itn, err := strconv.Atoi(it)
			if err != nil {
				panic(err)
			}
			lnn = append(lnn, itn)
		}

		in = append(in, lnn)
	}

	fmt.Printf("%v\n", part1(in))
	fmt.Printf("%v\n", part2(in))
}
