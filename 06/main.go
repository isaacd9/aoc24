package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func findStart(in grid) (int, int) {
	for i, row := range in {
		for j, ch := range row {
			if ch == "^" {
				return i, j
			}
		}
	}

	return -1, -1
}

type pos struct {
	x int
	y int
}

func inner(in grid) (map[pos]int, bool) {
	row, col := findStart(in)

	m := map[pos]int{
		{row, col}: 1,
	}

	visited := map[pos]string{}

	for {
		// fmt.Println(row, col, in[row][col], len(m))
		if row < 0 || row >= len(in) || col < 0 || col >= len(in[row]) {
			return m, false
		}
		v := in[row][col]

		// in.print()

		switch v {
		case "^": // up
			if row-1 < 0 {
				return m, false
			}

			if in[row-1][col] == "#" {
				in[row][col] = ">"
			} else {
				if visited[pos{row, col}] == v {
					return m, true
				}

				visited[pos{row, col}] = in[row][col]
				in[row-1][col] = "^"
				in[row][col] = "X"
				row--
			}
		case "v": // down
			if row+1 >= len(in) {
				return m, false
			}

			if in[row+1][col] == "#" {
				in[row][col] = "<"
			} else {
				if visited[pos{row, col}] == v {
					return m, true
				}
				visited[pos{row, col}] = in[row][col]
				in[row+1][col] = "v"
				in[row][col] = "X"
				row++
			}
		case "<": // left
			if col-1 < 0 {
				return m, false
			}

			if in[row][col-1] == "#" {
				in[row][col] = "^"
			} else {
				if visited[pos{row, col}] == v {
					return m, true
				}
				visited[pos{row, col}] = in[row][col]
				in[row][col-1] = "<"
				in[row][col] = "X"
				col--
			}
		case ">": // right
			if col+1 >= len(in[col]) {
				return m, false
			}

			if in[row][col+1] == "#" {
				in[row][col] = "v"
			} else {
				if visited[pos{row, col}] == v {
					return m, true
				}
				visited[pos{row, col}] = in[row][col]
				in[row][col+1] = ">"
				in[row][col] = "X"
				col++
			}
		}
		m[pos{row, col}] += 1
	}
}

func part1(in grid) int {
	m, _ := inner(in)
	return len(m)
}

func part2(in grid) (c int) {
	for i, row := range in {
		for j := range row {
			grid := in.copy()
			grid[i][j] = "#"
			if grid[i][j] == "^" {
				continue
			}
			// fmt.Println()
			// grid.print()
			_, loop := inner(grid)
			if loop {
				c++
			}
		}
	}
	return c
}

type grid [][]string

func (g grid) print() {
	for _, row := range g {
		fmt.Println(row)
	}
}

func (g grid) copy() grid {
	out := make(grid, len(g))
	for i, row := range g {
		out[i] = make([]string, len(row))
		copy(out[i], row)
	}
	return out
}

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var in grid

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
	fmt.Println(part1(in.copy()))
	fmt.Println(part2(in.copy()))
}
