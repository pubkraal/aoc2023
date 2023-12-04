package set

func Intersect[T comparable](a, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]bool)

	for _, v := range a {
		hash[v] = true
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
