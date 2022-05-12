package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}

	res := (Fibonacci(index-2) + Fibonacci(index-1))
	if res < 0 {
		return -res
	}
	return res
}
