package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
)

func main() {
	var t1 time.Duration
	var t2 time.Duration

	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	data := force.Must(input.ReadFileToStringSlice(filename))

	grid := makeGrid(len(data), len(data[0]), '.')
	for r, line := range data {
		for c, ch := range line {
			grid[r][c] = ch
		}
	}

	s1 := time.Now()
	p1 := calcWeight(roll(grid))
	t1 = time.Since(s1)

	s2 := time.Now()
	hist := make(map[string]int)
	target := 1000000000
	for i := 0; i < target; i++ {
		for j := 0; j < 4; j++ {
			grid = rotate(roll(grid))
		}
		h := flatten(grid)
		if v, ok := hist[h]; ok {
			cl := i - v
			amt := (target - i) / cl
			i += amt * cl
		}
		hist[h] = i
	}
	p2 := calcWeight(grid)
	t2 = time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func rotate(grid [][]rune) [][]rune {
	R := len(grid)
	C := len(grid[0])
	ng := makeGrid(R, C, '?')
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			ng[c][R-1-r] = grid[r][c]
		}
	}
	return ng
}

func roll(grid [][]rune) [][]rune {
	rows := len(grid)
	cols := len(grid[0])

	ng := makeGrid(rows, cols, '?')

	for c := 0; c < cols; c++ {
		for i := 0; i < rows; i++ {
			for r := 0; r < rows; r++ {
				if ng[r][c] == '?' {
					ng[r][c] = grid[r][c]
				}
				if ng[r][c] == 'O' && r > 0 && ng[r-1][c] == '.' {
					ng[r][c] = '.'
					ng[r-1][c] = 'O'
				}
			}
		}
	}

	return ng
}

func calcWeight(grid [][]rune) int {
	ret := 0
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'O' {
				ret += rows - r
			}
		}
	}
	return ret
}

func makeGrid(r, c int, def rune) [][]rune {
	rows := make([][]rune, r)
	for i := 0; i < r; i++ {
		rows[i] = make([]rune, c)
		for j := 0; j < c; j++ {
			rows[i][j] = def
		}
	}
	return rows
}

func flatten(grid [][]rune) string {
	str := ""
	for _, line := range grid {
		str += string(line)
	}
	return str
}
