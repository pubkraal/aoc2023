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

type gp struct {
	y int
	x int
}

type direction int32

const (
	Up direction = iota
	Right
	Down
	Left
)

type Instr struct {
	Dir   direction
	Steps int
	Color string
}

func main() {
	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)
	data := force.Must(input.ReadFileToStringSlice(filename))
	instrs := make([]Instr, len(data))
	for idx, line := range data {
		parts := strings.Split(line, " ")
		var dir direction
		switch parts[0] {
		case "U":
			dir = Up
		case "R":
			dir = Right
		case "D":
			dir = Down
		case "L":
			dir = Left
		}
		color := strings.Trim(parts[2], "()")
		instrs[idx] = Instr{dir, force.Must(strconv.Atoi(parts[1])), color}
	}

	dirs := make(map[direction][2]int)
	dirs[Up] = [2]int{-1, 0}
	dirs[Down] = [2]int{1, 0}
	dirs[Left] = [2]int{0, -1}
	dirs[Right] = [2]int{0, 1}
	dt := map[string]direction{
		"0": Right,
		"1": Down,
		"2": Left,
		"3": Up,
	}

	s1 := time.Now()
	cur := gp{0, 0}
	coords := make([]gp, 0)
	coords = append(coords, cur)
	length := 0
	for _, i := range instrs {
		m := dirs[i.Dir]
		length += i.Steps

		cur.y += m[0] * i.Steps
		cur.x += m[1] * i.Steps
		coords = append(coords, cur)
	}
	p1 := shoelace(coords) + (length / 2) + 1
	t1 := time.Since(s1)

	s2 := time.Now()
	cur = gp{0, 0}
	coords2 := make([]gp, 0)
	coords2 = append(coords2, cur)
	length = 0
	for _, i := range instrs {
		color := i.Color[1:6]
		dir := dt[i.Color[len(i.Color)-1:]]
		m := dirs[dir]
		s := int(force.Must(strconv.ParseInt(color, 16, 32)))

		cur.y += m[0] * s
		cur.x += m[1] * s
		length += s
		coords2 = append(coords2, cur)
	}
	p2 := shoelace(coords2) + (length / 2) + 1
	t2 := time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func shoelace(c []gp) int {
	n := len(c)
	s1 := 0
	s2 := 0

	for i := 0; i < n-1; i++ {
		s1 += c[i].x*c[i+1].y + 1
		s2 += c[i].y*c[i+1].x + 1
	}

	s1 += c[n-1].x * c[0].y
	s2 += c[0].x * c[n-1].y

	return abs(s1-s2) / 2
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
