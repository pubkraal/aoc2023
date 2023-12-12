package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/geo"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/itertools"
	"github.com/pubkraal/aoc2023/internal/util"
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
	data := force.Must(input.ReadFileToStringSlice(filename))

	p1 := 0
	p2 := 0

	s1 := time.Now()
	s2 := time.Now()

	er := make([]int, 0)
	ec := make([]int, 0)

	galaxies := make([]geo.Coord, 0)

	// Find empty rows and columns
	for ir, line := range data {
		cnt := itertools.NewCounter([]rune(line))
		if len(cnt) == 1 {
			er = append(er, ir)
		}
	}

	for ic := 0; ic < len(data[0]); ic++ {
		found := false

		for ir, line := range data {
			if line[ic] == '#' {
				galaxies = append(galaxies, geo.Coord{X: ic, Y: ir})
				found = true
			}
		}

		if !found {
			ec = append(ec, ic)
		}
	}

	p1Add := 1
	p2Add := 1000000 - 1

	done := make(map[geo.Coord]bool)
	for i := 0; i < len(galaxies); i++ {
		done[galaxies[i]] = true
		for j := 0; j < len(galaxies); j++ {
			if done[galaxies[j]] || i == j {
				continue
			}
			d1 := geo.ManhattanDistance(galaxies[i], galaxies[j])
			d2 := d1
			minr := util.Min(galaxies[i].Y, galaxies[j].Y)
			maxr := util.Max(galaxies[i].Y, galaxies[j].Y)
			minc := util.Min(galaxies[i].X, galaxies[j].X)
			maxc := util.Max(galaxies[i].X, galaxies[j].X)
			for _, r := range er {
				if minr < r && r < maxr {
					d1 += p1Add
					d2 += p2Add
				}
			}
			for _, c := range ec {
				if minc < c && c < maxc {
					d1 += p1Add
					d2 += p2Add
				}
			}
			p1 += d1
			p2 += d2
		}
	}

	t1 = time.Since(s1)
	t2 = time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}
