package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
)

type block struct {
	Src int
	Dst int
	Rng int
}

func main() {
	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	input := force.Must(input.ReadFileToStringSlice(filename))

	seeds := make([]int, 0)
	for _, seed := range strings.Split(input[0][7:], " ") {
		seeds = append(seeds, force.Must(strconv.Atoi(seed)))
	}

	blocks := make([][]block, 0)
	curblock := make([]block, 0)
	max := len(input) - 1

	for idx, line := range input[2:] {
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			blocks = append(blocks, curblock)
			curblock = make([]block, 0)
			continue
		}

		segments := strings.Split(line, " ")
		if len(segments) != 3 {
			panic("Invalid line: " + line)
		}

		curblock = append(curblock, block{
			force.Must(strconv.Atoi(segments[1])),
			force.Must(strconv.Atoi(segments[0])),
			force.Must(strconv.Atoi(segments[2])),
		})

		if (idx + 2) == max {
			blocks = append(blocks, curblock)
		}
	}

	blocks = blocks[1:]

	p1 := part1(seeds, blocks)
	p2 := part2(seeds, blocks)

	fmt.Println("01:", p1)
	fmt.Println("02:", p2)
}

func getPos(in []block, src int) int {
	for _, set := range in {
		if src >= set.Src && src < set.Src+set.Rng {
			return set.Dst + (src - set.Src)
		}
	}

	return src
}

func getLocation(seed int, blocks [][]block) int {
	res := seed

	for _, m := range blocks {
		res = getPos(m, res)
	}

	return res
}

func part1(seeds []int, blocks [][]block) int {
	lowest := -1

	for _, seed := range seeds {
		pos := getLocation(seed, blocks)

		if lowest == -1 || pos < lowest {
			lowest = pos
		}
	}
	return lowest
}

func part2(seeds []int, blocks [][]block) int {
	lowest := -1

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		rng := seeds[i+1]
		for x := 0; x < rng; x++ {
			curseed := start + x
			pos := getLocation(curseed, blocks)
			if lowest == -1 || pos < lowest {
				lowest = pos
			}
		}
	}

	return lowest
}
