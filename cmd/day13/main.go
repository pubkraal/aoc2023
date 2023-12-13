package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	data := force.Must(input.ReadFileToString(filename))
	blocks := strings.Split(data, "\n\n")

	s1 := time.Now()

	p1 := 0
	p2 := 0

part1:
	for _, block := range blocks {
		// Find vertical mirrors
		rb := rotate(block)
		lines := strings.Split(rb, "\n")
		for idx, line := range lines {
			if idx >= 1 && line == lines[idx-1] {
				if !mirror(lines, idx) {
					continue
				}
				p1 += idx
				continue part1
			}
		}

		// Find horizontal mirrors
		lines = strings.Split(block, "\n")
		for idx, line := range lines {
			if idx >= 1 && line == lines[idx-1] {
				if !mirror(lines, idx) {
					continue
				}
				p1 += 100 * idx
				break
			}
		}
	}
	t1 = time.Since(s1)

	s2 := time.Now()
part2:
	for _, block := range blocks {
		rb := rotate(block)
		lines := strings.Split(rb, "\n")
		for idx := 0; idx < len(lines); idx++ {
			if idx == 0 {
				continue
			}
			s := smidges(lines, idx)
			if s != 1 {
				continue
			} else {
				p2 += idx
				continue part2
			}
		}

		lines = strings.Split(block, "\n")
		for idx := 0; idx < len(lines); idx++ {
			if idx == 0 {
				continue
			}

			s := smidges(lines, idx)
			if s != 1 {
				continue
			} else {
				p2 += 100 * idx
				break
			}
		}
	}
	t2 = time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func rotate(in string) string {
	lines := strings.Split(in, "\n")
	nl := make([]string, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		s := ""
		for j := 0; j < len(lines); j++ {
			s += string(lines[j][i])
		}
		nl[i] = s
	}

	return strings.Join(nl, "\n")
}

func mirror(lines []string, idx int) bool {
	for i := 0; ; i++ {
		up := idx + i
		do := idx - (1 + i)
		if up >= len(lines) {
			break
		}
		if do < 0 {
			break
		}
		if lines[up] != lines[do] {
			return false
		}
	}

	return true
}

func smidges(lines []string, idx int) int {
	diffs := 0
	for i := 0; ; i++ {
		up := idx + i
		do := idx - (1 + i)
		if up >= len(lines) {
			break
		}
		if do < 0 {
			break
		}
		l1 := lines[up]
		l2 := lines[do]
		if len(l1) != len(l2) {
			panic("no")
		}
		for j := 0; j < len(l1); j++ {
			if l1[j] != l2[j] {
				diffs++
			}
		}
		// break early, don't bother anymore if more than 1
		if diffs > 1 {
			return diffs
		}
	}
	return diffs
}
