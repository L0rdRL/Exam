package piscine

func BasicAtoi2(s string) int {
	tx := []rune(s)
	var num int
	if s == "" {
		return 0
	}
	for i := 0; i < len(tx); i++ {
		if tx[i] < 47 || tx[i] > 57 {
			return 0
		}
		num *= 10
		num += int(tx[i] - 48)

	}
	return num
}
