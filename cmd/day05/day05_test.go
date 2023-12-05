package main

import "testing"

func TestGetPos(t *testing.T) {
	testblock := []block{
		{50, 98, 2},
		{52, 50, 48},
	}
	tables := []struct {
		blocks []block
		pos    int
		expect int
	}{
		{testblock, 10, 10},
		{testblock, 50, 98},
		{testblock, 51, 99},
		{testblock, 52, 50},
	}

	for _, tc := range tables {
		t.Run("", func(t *testing.T) {
			got := getPos(tc.blocks, tc.pos)
			if got != tc.expect {
				t.Errorf("getPos(%v, %v) = %v, want %v", tc.blocks, tc.pos, got, tc.expect)
			}
		})
	}
}
