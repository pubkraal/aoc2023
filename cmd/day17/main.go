package main

import (
	"container/heap"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
	"github.com/pubkraal/aoc2023/internal/util"
)

type Edge struct {
	Value int
	Row   int
	Col   int
	Dir   int
	Indir int
}

func (e Edge) H() HistoricEdge {
	return HistoricEdge{e.Row, e.Col, e.Dir, e.Indir}
}

// HistoricEdge is a very long name for an Edge with basic information that I'm
// gonna need later and more and everything
type HistoricEdge struct {
	Row   int
	Col   int
	Dir   int
	Indir int
}

type Edges []*Edge

func (es Edges) Len() int {
	return len(es)
}

func (es Edges) Less(i, j int) bool {
	return es[i].Value < es[j].Value
}

func (es Edges) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es *Edges) Push(x any) {
	item := x.(*Edge)
	*es = append(*es, item)
}

func (es *Edges) Pop() any {
	old := *es
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*es = old[0 : n-1]
	return item
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
	grid := make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		l := make([]int, len(data[i]))
		grid[i] = l
		for j := 0; j < len(data[0]); j++ {
			grid[i][j] = force.Must(strconv.Atoi(string(data[i][j])))
		}
	}

	s1 := time.Now()
	p1 := itsCalledDIJKstra(grid, false)
	t1 := time.Since(s1)

	s2 := time.Now()
	p2 := itsCalledDIJKstra(grid, true)
	t2 := time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func itsCalledDIJKstra(grid [][]int, part2 bool) int {
	rows := len(grid)
	cols := len(grid[0])
	q := make(Edges, 1)
	q[0] = &Edge{0, 0, 0, -1, -1}
	heap.Init(&q)

	done := make(map[HistoricEdge]int)

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for q.Len() > 0 {
		e := heap.Pop(&q).(*Edge)
		h := e.H()
		if _, ok := done[h]; ok {
			continue
		}

		done[h] = e.Value

		for i, dir := range directions {
			rr := e.Row + dir[0]
			cc := e.Col + dir[1]
			new_dir := i
			new_indir := e.Indir + 1
			if new_dir != e.Dir {
				new_indir = 1
			}

			isntreverse := (new_dir+2)%4 != e.Dir

			valid := false
			if !part2 {
				valid = new_indir <= 3
			} else {
				valid = new_indir <= 10 && (new_dir == e.Dir || e.Indir >= 4 || e.Indir == -1)
			}

			if 0 <= rr && rr < rows && 0 <= cc && cc < cols && isntreverse && valid {
				cost := grid[rr][cc]
				ne := &Edge{
					e.Value + cost,
					rr,
					cc,
					new_dir,
					new_indir,
				}
				heap.Push(&q, ne)
			}
		}
	}

	res := math.MaxInt
	for k, v := range done {
		if k.Row == rows-1 && k.Col == cols-1 {
			res = util.Min(res, v)
		}
	}

	return res
}
