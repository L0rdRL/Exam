package piscine

func IsLower(s string) bool {
	ss := []rune(s)
	var sss int

	for i := 0; i < len(ss); i++ {
		if ss[i] >= 97 && ss[i] <= 122 {
			sss++
		}
	}
	if sss == len(s) {
		return true
	} else {
		return false
	}
}
