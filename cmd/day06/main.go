package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

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

	times := util.Filter(strings.Split(input[0], " "))[1:]
	distances := util.Filter(strings.Split(input[1], " "))[1:]
	p2time := strings.Join(times, "")
	p2distance := strings.Join(distances, "")

	s1 := time.Now()
	p1 := calcRace(times, distances)
	t1 := time.Since(s1)
	s2 := time.Now()
	p2 := calcRace([]string{p2time}, []string{p2distance})
	t2 := time.Since(s2)

	fmt.Printf("01: %d (%v)\n", p1, t1)
	fmt.Printf("02: %d (%v)\n", p2, t2)
}

func calcRace(times, dinstances []string) int {
	res := 1
	for i := 0; i < len(times); i++ {
		time := force.Must(strconv.Atoi(times[i]))
		distance := force.Must(strconv.Atoi(dinstances[i]))

		term := math.Sqrt(math.Pow(float64(time), 2.0) - (4.0 * float64(distance)))
		lower := (float64(time) - term) / 2.0
		upper := (float64(time) + term) / 2.0
		res *= int(math.Floor(upper)) - int(math.Ceil(lower)) + 1
	}
	return res
}
