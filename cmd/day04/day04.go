package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/set"
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

	p1 := 0
	p2 := 0

	wincards := make(map[int]int)

	for idx, line := range input {
		wincards[idx] += 1
		var id int
		ident, game, _ := strings.Cut(line, ": ")
		card, numbers, _ := strings.Cut(game, " | ")
		fmt.Sscanf(ident, "Card %d", &id)
		c := util.SplitAndFilter(card, " ")
		n := util.SplitAndFilter(numbers, " ")

		match := set.Intersect(c, n)
		if len(match) > 0 {
			p1 += (1 << (len(match) - 1))
		}

		for x := 0; x < len(match); x++ {
			wincards[idx+1+x] += wincards[idx]
		}
	}

	for _, val := range wincards {
		p2 += val
	}

	fmt.Println("01:", p1)
	fmt.Println("02:", p2)
}
