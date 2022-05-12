package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ss := os.Args[1:]

	for i := 0; i < len(ss); i++ {
		sd := ss[i]
		for j := 0; j < len(sd); j++ {
			z01.PrintRune(rune(sd[j]))
		}
		z01.PrintRune('\n')
	}
}
