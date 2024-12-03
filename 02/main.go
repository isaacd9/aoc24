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
	rows := make([][]int, 0)

	for i := 0; i < len(row); i++ {
		var newRow []int
		for j, it := range row {
			if j == i {
				continue
			}
			newRow = append(newRow, it)
		}
		rows = append(rows, newRow)
	}

	log.Print("rows: ", rows)

	for _, r := range rows {
		log.Printf("r: %v", r)
		if dsc(r) || asc(r) {
			return true
		}
	}

	return false
}

func dsc2(row []int) bool {
	rows := make([][]int, 0)

	for i := 0; i < len(row); i++ {
		var newRow []int
		for j, it := range row {
			if j == i {
				continue
			}
			newRow = append(newRow, it)
		}
		rows = append(rows, newRow)
	}

	log.Print("rows: ", rows)
	for _, r := range rows {
		if dsc(r) || asc(r) {
			return true
		}
	}

	return false
}

func part2(in [][]int) int {
	var c int

	for _, row := range in {
		log.Printf("row: %v", row)
		if row[1] > row[0] || row[2] > row[0] {
			if asc2(row) {
				fmt.Println("ok")
				c++
			} else {
				fmt.Println("not ok")
			}
		} else if row[1] < row[0] || row[2] < row[0] {
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
