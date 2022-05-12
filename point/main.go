package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	for _, i := range "x = " {
		z01.PrintRune(i)
	}
	printNbr(points.x)
	z01.PrintRune(',')
	for _, i := range " y = " {
		z01.PrintRune(i)
	}
	printNbr(points.y)
	z01.PrintRune('\n')
}

func printNbr(a int) {
	var b []int
	if a == 0 {
		z01.PrintRune('0')
	}
	for a > 0 {
		b = append(b, a%10)
		a = a / 10
	}

	for i := len(b) - 1; i >= 0; i-- {
		z01.PrintRune(rune(b[i] + 48))
	}
}
