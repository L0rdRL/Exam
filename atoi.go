package piscine

func Atoi(s string) int {
	tx := []rune(s)
	var num int
	if s == "" {
		return 0
	}
	T := len(tx) - 1
	i := 0
	if tx[0] == 45 || tx[0] == 43 {
		i++
	}
	for ; i <= T; i++ {

		if tx[i] < 48 || tx[i] > 57 {
			return 0
		}
		num *= 10
		num += int(tx[i] - 48)
	}
	if tx[0] == 45 {
		return -num
	} else {
		return num
	}
}
