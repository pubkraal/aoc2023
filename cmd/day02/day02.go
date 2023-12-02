package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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

	p1 := 0
	p2 := 0
	for _, line := range input {
		game, data, _ := strings.Cut(line, ":")
		possible, power := gamePossible(data, 12, 13, 14)
		if possible {
			var num int
			fmt.Sscanf(game, "Game %d", &num)
			p1 += num
		}
		p2 += power
	}

	fmt.Println("01:", p1)
	fmt.Println("02:", p2)
}

func gamePossible(line string, red, green, blue int) (bool, int) {
	sets := strings.Split(line, ";")
	possible := true
	maxg := 0
	maxr := 0
	maxb := 0
	for _, set := range sets {
		colors := strings.Split(set, ",")
		g := 0
		r := 0
		b := 0
		for _, color := range colors {
			var count int
			var name string
			fmt.Sscanf(color, "%d %s", &count, &name)
			switch name {
			case "green":
				g += count
			case "red":
				r += count
			case "blue":
				b += count
			}
		}
		if g > green || r > red || b > blue {
			possible = false
		}
		maxg = util.Max(maxg, g)
		maxr = util.Max(maxr, r)
		maxb = util.Max(maxb, b)
	}

	return possible, (maxg * maxr * maxb)
}
