package itertools

func NewCounter[T comparable](a []T) map[T]bool {
	h := make(map[T]bool)

	for _, v := range a {
		h[v] = true
	}
	return h
}
