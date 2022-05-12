package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) >= 3 {

		fileNames := args[2:]

		for i := 0; i < len(fileNames); i++ {
			file, err := os.OpenFile(fileNames[i], os.O_RDWR, 0o644)

			if file != nil {
				fmt.Printf("==> " + fileNames[i] + " <==" + "\n")
			}

			check(err)
			defer file.Close()

			fmt.Printf(string(file[len(file)-atoi(args[1]):]) + "\n")
		}
	}
}

func atoi(input string) int {
	numbers := []byte(input)
	var result int

	for i := 0; i < len(numbers); i++ {
		if numbers[i] >= 48 && numbers[i] <= 57 {
			result *= 10
			result += int(numbers[i]) - 48
		}
	}

	return result
}

func check(err error) {
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(1)
	}
}
