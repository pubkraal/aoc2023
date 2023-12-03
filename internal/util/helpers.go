package util

import "sort"

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
