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

	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	data := force.Must(input.ReadFileToStringSlice(filename))

	p1 := 0
	p2 := 0

	s1 := time.Now()
	for _, line := range data {
		nums := force.Must(input.StringToIntSlice(line))
		pre, post := predict(nums)
		p1 += post
		p2 += pre
	}
	t1 = time.Since(s1)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t1)
}

func predict(line []int) (int, int) {
	if isAllZero(line) {
		return 0, 0
	}
	diffs := make([]int, len(line)-1)

	for i := 1; i < len(line); i++ {
		diffs[i-1] = line[i] - line[i-1]
	}

	pre, post := predict(diffs)
	return line[0] - pre, line[len(line)-1] + post
}

func isAllZero(line []int) bool {
	for _, num := range line {
		if num != 0 {
			return false
		}
	}

	return true
}
