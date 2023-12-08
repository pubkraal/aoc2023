package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/math"
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
	input := force.Must(input.ReadFileToStringSlice(filename))

	directions := input[0]
	routes := make(map[string][]string)

	for _, instr := range input[2:] {
		var start string
		var left string
		var right string

		fmt.Sscanf(instr, "%3s = (%3s, %3s)", &start, &left, &right)

		routes[start] = []string{left, right}
	}

	currents := make([]string, 0)
	for key := range routes {
		if strings.HasSuffix(key, "A") {
			currents = append(currents, key)
		}
	}

	p1 := 0
	p2 := 0

	s1 := time.Now()

	move := func(cur string, counter int) string {
		dir := directions[counter%len(directions)]
		var new string
		if dir == 'L' {
			new = routes[cur][0]
		} else {
			new = routes[cur][1]
		}
		return new
	}

	current := "AAA"
	for {
		current = move(current, p1)
		p1++

		if current == "ZZZ" {
			break
		}
	}

	t1 = time.Since(s1)

	counts := make([]int, len(currents))

	s2 := time.Now()
	for i := 0; i < len(currents); i++ {
		cnt := 0
		cur := currents[i]
		for {
			cur = move(cur, cnt)
			cnt++

			if strings.HasSuffix(cur, "Z") {
				break
			}
		}
		counts[i] = cnt
	}
	p2 = math.LCMOfSlice(counts)
	t2 = time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}
