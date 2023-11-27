package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pubkraal/aoc2023/internal/input"
)

func main() {
	fmt.Println("Day 1")

	// Define flags
	filename := flag.String("file", "", "Input filename")
	flag.Parse()

	// Check if filename is provided
	if *filename == "" {
		fmt.Println("Please provide an input filename")
		os.Exit(1)
	}

	// Read input file into string
	input := input.ReadFileToString(*filename)

	fmt.Println(input)
}
