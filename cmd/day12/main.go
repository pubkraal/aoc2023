package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
)

type Key struct {
	PosDot   int
	PosBlock int
	Current  int
}

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

	s1 := time.Now()
	s2 := time.Now()

	p1 := 0
	p2 := 0

	for _, part2 := range []bool{false, true} {
		for _, line := range data {
			dots, blockstxt, _ := strings.Cut(line, " ")

			if part2 {
				dots = strings.Join([]string{dots, dots, dots, dots, dots}, "?")
				blockstxt = strings.Join([]string{blockstxt, blockstxt, blockstxt, blockstxt, blockstxt}, ",")
			}

			blockssplit := strings.Split(blockstxt, ",")
			blocks := make([]int, len(blockssplit))
			for i := 0; i < len(blockssplit); i++ {
				blocks[i] = force.Must(strconv.Atoi(blockssplit[i]))
			}

			cache := make(map[Key]int)
			score := process(cache, dots, blocks, 0, 0, 0)
			if !part2 {
				p1 += score
			} else {
				p2 += score
			}
		}

		if !part2 {
			t1 = time.Since(s1)
		} else {
			t2 = time.Since(s2)
		}
	}

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func process(cache map[Key]int, dots string, blocks []int, posDot, posBlock, current int) int {
	key := Key{PosDot: posDot, PosBlock: posBlock, Current: current}
	if val, ok := cache[key]; ok {
		return val
	}
	if posDot == len(dots) {
		if posBlock == len(blocks) && current == 0 {
			return 1
		} else if posBlock == len(blocks)-1 && blocks[posBlock] == current {
			return 1
		} else {
			return 0
		}
	}

	ans := 0
	for _, b := range []byte{'.', '#'} {
		if dots[posDot] == b || dots[posDot] == '?' {
			if b == '.' && current == 0 {
				ans += process(cache, dots, blocks, posDot+1, posBlock, 0)
			} else if b == '.' && current > 0 && posBlock < len(blocks) && blocks[posBlock] == current {
				ans += process(cache, dots, blocks, posDot+1, posBlock+1, 0)
			} else if b == '#' {
				ans += process(cache, dots, blocks, posDot+1, posBlock, current+1)
			}
		}
	}

	cache[key] = ans
	return ans
}
