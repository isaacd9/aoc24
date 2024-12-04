package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func get(in [][]string, i, j int) string {
	if i < 0 || j < 0 || i >= len(in) || j >= len(in[i]) {
		return ""
	}
	return in[i][j]
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func check1(in [][]string, i, j int) (c int) {
	/*
		left := get(in, i, j)+get(in, i, j-1)+get(in, i, j-2)+get(in, i, j-3) == "XMAS"
		top := get(in, i, j)+get(in, i+1, j)+get(in, i+2, j)+get(in, i+3, j) == "XMAS"
		right := get(in, i, j)+get(in, i, j+1)+get(in, i, j+2)+get(in, i, j+3) == "XMAS"
		bottom := get(in, i, j)+get(in, i-1, j)+get(in, i-2, j)+get(in, i-3, j) == "XMAS"

		bleft := get(in, i, j)+get(in, i, j-1)+get(in, i, j-2)+get(in, i, j-3) == "SAMX"
		btop := get(in, i, j)+get(in, i+1, j)+get(in, i+2, j)+get(in, i+3, j) == "SAMX"
		bright := get(in, i, j)+get(in, i, j+1)+get(in, i, j+2)+get(in, i, j+3) == "SAMX"
		bbottom := get(in, i, j)+get(in, i-1, j)+get(in, i-2, j)+get(in, i-3, j) == "SAMX"

		return btoi(left) + btoi(top) + btoi(right) + btoi(bottom) + btoi(bleft) + btoi(btop) + btoi(bright) + btoi(bbottom)
	*/

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			var a string
			for c := 0; c < 4; c++ {
				a += get(in, i+x*c, j+y*c)
			}
			log.Println(a)
			if a == "XMAS" {
				c += 1
			}
		}
	}

	return c
}

func part1(in [][]string) (c int) {
	for i := range in {
		for j := range in[i] {
			c += check1(in, i, j)
		}
	}

	return c
}

func check2(in [][]string, i, j int) (c int) {
	if get(in, i, j) != "A" {
		return 0
	}

	fst := get(in, i-1, j-1) + get(in, i, j) + get(in, i+1, j+1)
	snd := get(in, i-1, j+1) + get(in, i, j) + get(in, i+1, j-1)

	log.Println(i, j, fst, snd)

	if fst == "MAS" && snd == "SAM" {
		log.Println("ok")
		return 1
	}

	if fst == "SAM" && snd == "MAS" {
		log.Println("ok")
		return 1
	}

	if fst == "MAS" && snd == "MAS" {
		log.Println("ok")
		return 1
	}

	if fst == "SAM" && snd == "SAM" {
		log.Println("ok")
		return 1
	}

	return 0
}

func part2(in [][]string) (c int) {
	for i := range in {
		for j := range in[i] {
			c += check2(in, i, j)
		}
	}

	return c
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var in [][]string

	for i, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}
		in = append(in, []string{})
		for _, ch := range line {
			in[i] = append(in[i], string(ch))
		}
	}

	fmt.Println(in)
	fmt.Println(part1(in))
	fmt.Println(part2(in))
}
