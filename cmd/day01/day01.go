package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
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

	p1 := 0
	p2 := 0

	for _, line := range input {
		p1digits := make([]int, 0)
		p2digits := make([]int, 0)
		for i, r := range line {
			if unicode.IsDigit(r) {
				p1digits = append(p1digits, int(r-'0'))
				p2digits = append(p2digits, int(r-'0'))
				continue
			}
			switch {
			case gotNamedDigit(line, i, "one"):
				p2digits = append(p2digits, 1)
			case gotNamedDigit(line, i, "two"):
				p2digits = append(p2digits, 2)
			case gotNamedDigit(line, i, "three"):
				p2digits = append(p2digits, 3)
			case gotNamedDigit(line, i, "four"):
				p2digits = append(p2digits, 4)
			case gotNamedDigit(line, i, "five"):
				p2digits = append(p2digits, 5)
			case gotNamedDigit(line, i, "six"):
				p2digits = append(p2digits, 6)
			case gotNamedDigit(line, i, "seven"):
				p2digits = append(p2digits, 7)
			case gotNamedDigit(line, i, "eight"):
				p2digits = append(p2digits, 8)
			case gotNamedDigit(line, i, "nine"):
				p2digits = append(p2digits, 9)
			}
		}
		p1 += (p1digits[0] * 10) + p1digits[len(p1digits)-1]
		p2 += (p2digits[0] * 10) + p2digits[len(p2digits)-1]
	}
	fmt.Println("01:", p1)
	fmt.Println("02:", p2)
}

func gotNamedDigit(line string, i int, name string) bool {
	stop := i + len(name)
	if len(line) >= stop && line[i:stop] == name {
		return true
	}
	return false
}
