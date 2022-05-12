package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		ss := []string(os.Args[1:])
		var s int

		if ss[0] == "--upper" {
			for i := 1; i < len(ss); i++ {
				s = Atoi(ss[i])
				if s > 0 && s <= 26 {
					z01.PrintRune(rune(s + 64))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for i := 0; i < len(ss); i++ {
				s = Atoi(ss[i])
				if s > 0 && s <= 26 {
					z01.PrintRune(rune(s + 96))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	sd := []rune(s)
	var numv int
	if s == "" {
		return 0
	}
	B := len(sd) - 1
	i := 0
	if sd[0] == 45 || sd[0] == 43 {
		i++
	}

	for ; i <= B; i++ {

		if sd[i] < 48 || sd[i] > 57 {
			return 0
		}
		numv *= 10
		numv += int(sd[i] - 48)

	}
	if sd[0] == 45 {
		return -numv
	} else {
		return numv
	}
}
