package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nn []int
	if n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {
		nn = append(nn, n%10)
		n /= 10
	}
	SortIntegerTable(nn)
	for i := 0; i < len(nn); i++ {
		z01.PrintRune(rune(nn[i] + '0'))
	}
}
