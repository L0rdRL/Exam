package main

import "github.com/01-edu/z01"

func main() {
	a := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i <= len(a)-1; i++ {
		z01.PrintRune([]rune(a)[i])
	}
	z01.PrintRune('\n')
}
