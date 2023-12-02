package util

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
