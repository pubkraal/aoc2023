package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/util"
)

func main() {
	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	input := force.Must(input.ReadFileToStringSlice(filename))

	p1 := part1(input)
	p2 := part2(input)

	fmt.Println("01:", p1)
	fmt.Println("02:", p2)
}

func part1(input []string) int {
	sum := 0
	// loop until you find a number
	for li, line := range input {
		for ci := 0; ci < len(line); ci++ {
			r := line[ci]
			if !unicode.IsDigit(rune(r)) {
				continue
			}

			// Find digit, iterate further right until end of number
			start := ci
			for ci+1 < len(line) && unicode.IsDigit(rune(line[ci+1])) {
				ci++
			}
			end := ci

			number := force.Must(strconv.Atoi(line[start : end+1]))

			// Scan around the number for a symbol
			// Scan left, right
			if start > 0 {
				if !unicode.IsDigit(rune(line[start-1])) && line[start-1] != '.' {
					sum += number
					continue
				}
			}
			if end < len(line)-1 {
				if !unicode.IsDigit(rune(line[end+1])) && line[end+1] != '.' {
					sum += number
					continue
				}
			}

			// Scan top line
			if li > 0 {
				for ix := util.Max(start-1, 0); ix <= util.Min(end+1, len(line)-1); ix++ {
					ir := input[li-1][ix]
					if !unicode.IsDigit(rune(ir)) && ir != '.' {
						sum += number
						continue
					}
				}
			}

			// Scan bottom line
			if li < len(input)-1 {
				for ix := util.Max(start-1, 0); ix <= util.Min(end+1, len(line)-1); ix++ {
					ir := input[li+1][ix]
					if !unicode.IsDigit(rune(ir)) && ir != '.' {
						sum += number
						continue
					}
				}
			}
		}
	}
	return sum
}

func part2(input []string) int {
	sum := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] != '*' {
				continue
			}

			found := make([]int, 0)

			// Search for 2 distinct numbers in the surrounding cells
			if x > 0 && unicode.IsDigit(rune(input[y][x-1])) {
				number, _ := findNumber(input[y], x-1)
				found = append(found, number)
			}
			if x < len(input[y])-1 && unicode.IsDigit(rune(input[y][x+1])) {
				number, _ := findNumber(input[y], x+1)
				found = append(found, number)
			}

			if y > 0 {
				for ix := util.Max(x-1, 0); ix <= util.Min(x+1, len(input[y])-1); ix++ {
					r := input[y-1][ix]
					if unicode.IsDigit(rune(r)) {
						number, pos := findNumber(input[y-1], ix)
						found = append(found, number)
						ix = pos[len(pos)-1] + 1
					}
				}
			}
			if y < len(input)-1 {
				for ix := util.Max(x-1, 0); ix <= util.Min(x+1, len(input[y])-1); ix++ {
					r := input[y+1][ix]
					if unicode.IsDigit(rune(r)) {
						number, pos := findNumber(input[y+1], ix)
						found = append(found, number)
						ix = pos[len(pos)-1] + 1
					}
				}
			}

			if len(found) == 2 {
				sum += found[0] * found[1]
			}
		}
	}

	return sum
}

func findNumber(input string, x int) (int, []int) {
	if !unicode.IsDigit(rune(input[x])) {
		return 0, []int{}
	}
	// Find start
	seen := make(map[int]bool)
	start := 0
	end := 0
	for i := x; i >= 0; i-- {
		seen[i] = true
		if i == 0 || !unicode.IsDigit(rune(input[i-1])) {
			start = i
			break
		}
	}

	for i := x; i < len(input); i++ {
		seen[i] = true
		if i == len(input)-1 || !unicode.IsDigit(rune(input[i+1])) {
			end = i
			break
		}
	}

	seenpos := util.GetSortedKeys(seen)

	number := force.Must(strconv.Atoi(input[start : end+1]))
	return number, seenpos
}
