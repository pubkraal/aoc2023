package main

import "testing"

func TestGamePossible(t *testing.T) {
	tables := []struct {
		Name     string
		Game     string
		MaxRed   int
		MaxGreen int
		MaxBlue  int
		Possible bool
		Power    int
	}{
		{"Game 1", "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 12, 13, 14, true, 48},
		{"Game 2", "1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12, 13, 14, true, 12},
		{"Game 3", "8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 12, 13, 14, false, 1560},
		{"Game 4", "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 12, 13, 14, false, 630},
		{"Game 5", "6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 12, 13, 14, true, 36},
	}

	for _, tc := range tables {
		t.Run(tc.Name, func(t *testing.T) {
			possible, pwr := gamePossible(tc.Game, tc.MaxRed, tc.MaxGreen, tc.MaxBlue)
			if possible != tc.Possible {
				t.Fatalf("expected %v, got %v", tc.Possible, possible)
			}
			if pwr != tc.Power {
				t.Fatalf("expected %v, got %v", tc.Power, pwr)
			}
		})
	}
}
