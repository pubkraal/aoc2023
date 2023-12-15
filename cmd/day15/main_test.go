package main

import "testing"

func TestTsja256(t *testing.T) {
	tables := []struct {
		in     string
		expect int
	}{
		{"HASH", 52},
		{"rn=1", 30},
		{"pc=4", 180},
	}

	for _, tc := range tables {
		t.Run(tc.in, func(t *testing.T) {
			res := tsja256(tc.in)
			if res != tc.expect {
				t.Fatalf("expected %v, got %v", tc.expect, res)
			}
		})
	}
}
