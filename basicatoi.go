package piscine

func BasicAtoi(s string) int {
	tx := []rune(s)
	var num int
	for i := 0; i < len(tx); i++ {
		num *= 10
		num += int(tx[i] - 48)
	}
	return num
}
