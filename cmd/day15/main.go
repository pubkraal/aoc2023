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
	"github.com/pubkraal/aoc2023/internal/util"
)

type bv struct {
	Name string
	Val  int
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

	s1 := time.Now()
	b := strings.Split(data[0], ",")
	for _, c := range b {
		p1 += tsja256(c)
	}
	t1 = time.Since(s1)
	s2 := time.Now()
	boxes := make([][]bv, 256)
	for _, c := range b {
		if strings.ContainsRune(c, '=') {

			name := c[:len(c)-2]
			idx := tsja256(name)

			l := force.Must(strconv.Atoi(string(c[len(c)-1])))
			ns := getNames(boxes[idx])
			if util.In(name, ns) {
				// do thing
				nbv := make([]bv, 0)
				for _, v := range boxes[idx] {
					nv := v
					if nv.Name == name {
						nv.Val = l
					}
					nbv = append(nbv, nv)
				}
				boxes[idx] = nbv
			} else {
				boxes[idx] = append(boxes[idx], bv{name, l})
			}

		} else if strings.HasSuffix(c, "-") {

			name := c[:len(c)-1]
			idx := tsja256(name)

			cbv := boxes[idx]
			if len(cbv) == 0 {
				continue
			}

			nbv := make([]bv, 0)
			for _, v := range cbv {
				if v.Name == name {
					continue
				}
				nbv = append(nbv, v)
			}

			boxes[idx] = nbv

		}
	}

	for i, l := range boxes {
		for j, b := range l {
			p2 += (i + 1) * (j + 1) * b.Val
		}
	}

	t2 = time.Since(s2)

	fmt.Printf("01: %d (%s)\n", p1, t1)
	fmt.Printf("02: %d (%s)\n", p2, t2)
}

func tsja256(in string) int {
	res := 0
	for _, c := range in {
		res = ((res + int(c)) * 17) % 256
	}

	return res
}

func getNames(box []bv) []string {
	n := make([]string, len(box))

	for i, x := range box {
		n[i] = x.Name
	}

	return n
}
