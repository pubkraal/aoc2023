package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	times := util.Filter(strings.Split(input[0], " "))[1:]
	distances := util.Filter(strings.Split(input[1], " "))[1:]
	p2time := strings.Join(times, "")
	p2distance := strings.Join(distances, "")

	p1 := calcRace(times, distances)
	p2 := calcRace([]string{p2time}, []string{p2distance})

	fmt.Println("01: ", p1)
	fmt.Println("02: ", p2)
}

func getDistanceTravelled(time, maxtime int) int {
	return time * (maxtime - time)
}

func calcRace(times, distances []string) int {
	res := 1
	for i := 0; i < len(times); i++ {
		cur_time := force.Must(strconv.Atoi(times[i]))
		cur_distance := force.Must(strconv.Atoi(distances[i]))

		c := cur_time

		for j := 0; j < cur_time; j++ {
			dist := getDistanceTravelled(j, cur_time)
			if dist > cur_distance {
				c = j
				break
			}
		}

		num := cur_time - (2*c - 1)
		res *= num
	}
	return res
}
