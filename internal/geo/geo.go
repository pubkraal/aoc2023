package geo

import "github.com/pubkraal/aoc2023/internal/math"

type Coord struct {
	X int
	Y int
}

func GetNeighbors(m []string, pos Coord) map[Coord]rune {
	res := make(map[Coord]rune)

	lookup := []Coord{
		{X: pos.X - 1, Y: pos.Y},
		{X: pos.X + 1, Y: pos.Y},
		{X: pos.X, Y: pos.Y - 1},
		{X: pos.X, Y: pos.Y + 1},
	}

	for l := range lookup {
		if lookup[l].Y < 0 || lookup[l].Y >= len(m) {
			continue
		}
		if lookup[l].X < 0 || lookup[l].X >= len(m[lookup[l].Y]) {
			continue
		}
		res[lookup[l]] = rune(m[lookup[l].Y][lookup[l].X])
	}

	return res
}

func GetDiagonalNeighbors(m []string, pos Coord) map[Coord]rune {
	result := make(map[Coord]rune)

	for y := pos.Y - 1; y <= pos.Y+1; y++ {
		for x := pos.X - 1; x <= pos.X+1; x++ {
			if y < 0 || y >= len(m) {
				continue
			}
			if x < 0 || x >= len(m[y]) {
				continue
			}
			if x == pos.X && y == pos.Y {
				continue
			}
			result[Coord{X: x, Y: y}] = rune(m[y][x])
		}
	}

	return result
}

func ManhattanDistance(a, b Coord) int {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}
