package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		argss, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			for _, c := range "ERROR: " {
				z01.PrintRune(rune(c))
			}
			for _, c := range string(err.Error()) {
				z01.PrintRune(c)
			}

			z01.PrintRune('\n')
		}
		os.Stdout.WriteString(string(argss))

	} else {
		argss := os.Args[1:]
		for _, c := range argss {
			PrintArg(c)
		}
	}
}

func PrintArg(args string) {
	a := string(args)
	file, err := ioutil.ReadFile(a)
	if err != nil {
		for _, c := range "ERROR: " {
			z01.PrintRune(rune(c))
		}
		for _, c := range string(err.Error()) {
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')
		os.Exit(1)

	}

	for _, c := range string(file) {
		z01.PrintRune(rune(c))
	}
}
