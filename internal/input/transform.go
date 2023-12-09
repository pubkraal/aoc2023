package input

import (
	"strconv"

	"github.com/pubkraal/aoc2023/internal/util"
)

func StringToIntSlice(input string) ([]int, error) {
	var err error
	values := util.SplitAndFilter(input, " ")
	ret := make([]int, len(values))
	for i, v := range values {
		ret[i], err = strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
	}

	return ret, nil
}
