package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	a := 0
	if n/10 == 0 {
		a = n
		if n < 0 {
			z01.PrintRune('-')
		}
	} else {
		a = n % 10
		PrintNbr(n / 10)
	}
	if a < 0 {
		a = -a
	}
	z01.PrintRune(rune(a + 48))
}
