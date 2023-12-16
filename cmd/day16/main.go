package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/util"
)

type GP struct {
	X int
	Y int
}

type Heading int

const (
	Up Heading = iota
	Right
	Down
	Left
)

type vector struct {
	X       int
	Y       int
	heading Heading
}

func (v vector) GP() GP {
	return GP{v.X, v.Y}
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

	p1 := 0
	p2 := 0

	it := make(map[GP]rune)
	for r, line := range data {
		for c, ru := range line {
			if ru == '.' {
				continue
			}
			it[GP{c, r}] = ru
		}
	}

	s1 := time.Now()
	p1 = lightItUp(it, data, vector{-1, 0, Right})
	t1 = time.Since(s1)

	s2 := time.Now()
	for i := 0; i < len(data); i++ {
		p2 = util.Max(p2, lightItUp(it, data, vector{-1, i, Right}))
		p2 = util.Max(p2, lightItUp(it, data, vector{len(data[0]), i, Left}))
	}
	for i := 0; i < len(data[0]); i++ {
		p2 = util.Max(p2, lightItUp(it, data, vector{-1, i, Down}))
		p2 = util.Max(p2, lightItUp(it, data, vector{len(data), i, Up}))
	}
	t2 = time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func show(been map[GP]bool) {
	// Find biggest
	x := 0
	y := 0
	for k := range been {
		x = util.Max(k.X, x)
		y = util.Max(k.Y, y)
	}

	for r := 0; r <= y; r++ {
		for c := 0; c <= x; c++ {
			if _, ok := been[GP{c, r}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func lightItUp(items map[GP]rune, data []string, start vector) int {
	been := make(map[GP]bool)
	hds := make([]vector, 1)
	hds[0] = start
	groundhog := make(map[vector]bool)

	for len(hds) > 0 {
		nhds := make([]vector, 0)
		for i := 0; i < len(hds); i++ {
			c := hds[i]
			if _, ok := groundhog[c]; ok {
				continue
			}
			groundhog[c] = true

			// Move all by one
			switch c.heading {
			case Up:
				c.Y -= 1
			case Down:
				c.Y += 1
			case Left:
				c.X -= 1
			case Right:
				c.X += 1
			}

			// If it falls off the board, delete it
			if c.Y < 0 || c.Y >= len(data) || c.X < 0 || c.X >= len(data[0]) {
				continue
			}

			// Record position in been
			been[c.GP()] = true

			if v, ok := items[GP{c.X, c.Y}]; ok {
				switch v {
				case '\\':
					switch c.heading {
					case Up:
						c.heading = Left
					case Down:
						c.heading = Right
					case Right:
						c.heading = Down
					case Left:
						c.heading = Up
					}

					nhds = append(nhds, c)
				case '/':
					switch c.heading {
					case Up:
						c.heading = Right
					case Down:
						c.heading = Left
					case Right:
						c.heading = Up
					case Left:
						c.heading = Down
					}
					nhds = append(nhds, c)
				case '|':
					if c.heading == Right || c.heading == Left {
						c2 := c
						c.heading = Up
						c2.heading = Down

						nhds = append(nhds, c, c2)
					} else {
						nhds = append(nhds, c)
					}
				case '-':
					if c.heading == Up || c.heading == Down {
						c2 := c
						c.heading = Left
						c2.heading = Right

						nhds = append(nhds, c, c2)
					} else {
						nhds = append(nhds, c)
					}
				}
			} else {
				// Continue as if nothing happened
				nhds = append(nhds, c)
			}
		}
		hds = nhds
	}

	if len(data) < 20 {
		show(been)
	}

	return len(been)

}
