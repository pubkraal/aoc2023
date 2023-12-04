package set

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInserect(t *testing.T) {
	tables := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
	}

	for idx, tc := range tables {
		t.Run(fmt.Sprintf("TestI)ntersect %d", idx), func(t *testing.T) {
			got := Intersect(tc.a, tc.b)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("Intersect(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expect)
			}
		})
	}
}
