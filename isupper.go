package piscine

func IsUpper(s string) bool {
	ss := []rune(s)
	var sss int

	for i := 0; i < len(ss); i++ {
		if ss[i] >= 65 && ss[i] <= 90 {
			sss++
		}
	}
	if sss == len(s) {
		return true
	} else {
		return false
	}
}
