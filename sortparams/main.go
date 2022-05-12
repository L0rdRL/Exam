package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ss := []string(os.Args[1:])
	for i := len(ss); i > 0; i-- {
		for j := 1; j < i; j++ {
			if ss[j-1] > ss[j] {
				ss[j-1], ss[j] = ss[j], ss[j-1]
			}
		}
	}
	for i := 0; i < len(ss); i++ {
		sd := ss[i]
		for j := 0; j < len(sd); j++ {
			z01.PrintRune(rune(sd[j]))
		}
		z01.PrintRune('\n')
	}
}
