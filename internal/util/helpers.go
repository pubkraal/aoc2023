package util

import (
	"slices"
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

func GetKeys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func GetSortedKeys[T Inty, V any](m map[T]V) []T {
	keys := GetKeys(m)

	slices.Sort(keys)
	return keys
}

func GetValues[T comparable, V any](m map[T]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
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

func In[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
