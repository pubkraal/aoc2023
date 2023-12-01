package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	namedDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range input {
		p1digits := make([]int, 0)
		p2digits := make([]int, 0)
		for i, r := range line {
			if unicode.IsDigit(r) {
				p1digits = append(p1digits, int(r-'0'))
				p2digits = append(p2digits, int(r-'0'))
				continue
			}

			for j, name := range namedDigits {
				if strings.HasPrefix(line[i:], name) {
					p2digits = append(p2digits, j+1)
				}
			}
		}
		p1 += (p1digits[0] * 10) + p1digits[len(p1digits)-1]
		p2 += (p2digits[0] * 10) + p2digits[len(p2digits)-1]
	}
	fmt.Println("01:", p1)
	fmt.Println("02:", p2)
}
