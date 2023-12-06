package util

import (
	"sort"
	"strings"
)

type Inty interface {
	int | int8 | int16 | int32 | int64
}

func Max[T Inty](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Inty](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func GetSortedKeys[T Inty](m map[T]bool) []T {
	keys := make([]T, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func Filter(a []string) []string {
	ret := make([]string, 0)
	for _, part := range a {
		if len(part) > 0 {
			ret = append(ret, part)
		}
	}
	return ret
}

func SplitAndFilter(s string, sep string) []string {
	parts := strings.Split(s, sep)
	ret := make([]string, 0)
	for _, part := range parts {
		if len(part) > 0 {
			ret = append(ret, part)
		}
	}
	return ret
}
