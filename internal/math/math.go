package math

func LCMOfSlice(numbers []int) int {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = LCM(result, num)
	}
	return result
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
