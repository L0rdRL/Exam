package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		if len(args) == 3 {
			if IsNumeric(args[0]) && IsNumeric(args[2]) {
				if len(args[0]) != len("9223372036854775809") && len(args[2]) != len("9223372036854775809") {
					first := Atoi(args[0])
					operator := args[1]
					second := Atoi(args[2])
					if first < 2147483647 && second < 2147483647 {
						if typeof(first) != "int64" && typeof(second) != "int64" {
							num, err := Math(first, second, operator)
							if err != "a" && err != "b" {
								fmt.Println(err)
							} else if err != "b" {
								fmt.Println(num)
							}
						}
					}
				}
			}
		}
	}
}

func Math(a, b int32, c string) (int32, string) {
	var num int32
	var err string
	switch c {
	case "+":
		num, err = a+b, "a"
	case "-":
		num, err = a-b, "a"
	case "/":
		if b == 0 {
			return 0, "No division by 0"
		}
		num, err = a/b, "a"
	case "*":
		num, err = a*b, "a"
	case "%":
		if b == 0 {
			return 0, "No modulo by 0"
		}
		num, err = a%b, "a"
	default:
		num, err = 0, "b"
	}
	return num, err
}

func Atoi(s string) int32 {
	sd := []rune(s)
	var numv int32
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
		numv += int32(sd[i] - 48)

	}
	if sd[0] == 45 {
		return -numv
	} else {
		return numv
	}
}

func IsNumeric(s string) bool {
	ss := []rune(s)
	i := 0
	if ss[0] == '-' {
		i++
	}
	for ; i < len(ss); i++ {
		if ss[i] < 47 || ss[i] > 57 {
			return false
		}
	}
	return true
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
