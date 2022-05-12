package piscine

func IsNumeric(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] < 47 || ss[i] > 57 {
			return false
		}
	}
	return true
}
