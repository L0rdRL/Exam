package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(s []string) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			z01.PrintRune(rune(s[i][j]))
		}
		z01.PrintRune(rune('\n'))
	}
}
