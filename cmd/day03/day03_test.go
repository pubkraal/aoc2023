package main

import (
	"reflect"
	"testing"
)

func TestFindNumber(t *testing.T) {
	tables := []struct {
		input     string
		start     int
		value     int
		positions []int
	}{
		{"..592.....", 3, 592, []int{2, 3, 4}},
		{"..35..633.", 3, 35, []int{2, 3}},
		{"..35..633.", 6, 633, []int{6, 7, 8}},
		{"..35..633.", 2, 35, []int{2, 3}},
		{"..35..633.", 8, 633, []int{6, 7, 8}},
		{"..35..633.", 9, 0, []int{}},
		{"617*......", 2, 617, []int{0, 1, 2}},
		{"........23", 8, 23, []int{8, 9}},
	}

	for _, tc := range tables {
		t.Run(tc.input, func(t *testing.T) {
			got, pos := findNumber(tc.input, tc.start)
			if got != tc.value {
				t.Errorf("findNumber(%s, %d) = %d, _, want %d", tc.input, tc.start, got, tc.value)
			}
			if !reflect.DeepEqual(pos, tc.positions) {
				t.Errorf("findNumber(%s, %d) = _, %v, want %v", tc.input, tc.start, pos, tc.positions)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	got := part2(input)
	if got != 467835 {
		t.Errorf("part2(input) = %d, want %d", got, 467835)
	}
}

func TestPart1(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	got := part1(input)
	if got != 4361 {
		t.Errorf("part2(input) = %d, want %d", got, 4361)
	}
}
