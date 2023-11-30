package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pubkraal/aoc2023/internal/force"
	"github.com/pubkraal/aoc2023/internal/input"
)

func main() {
	// Define flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filename := flag.Arg(0)

	// Read input file into string
	input := force.Must(input.ReadFileToString(filename))

	fmt.Printf(input)
}
