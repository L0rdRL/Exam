package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		args := os.Args[1]
		a := string(args)
		file, err := ioutil.ReadFile(a)
		if err != nil {
			a := (string(err.Error()))
			fmt.Println(a)
		}
		fmt.Printf(string(file))

	}
}
