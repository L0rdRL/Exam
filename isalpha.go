package piscine

func IsAlpha(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] > 0 && ss[i] < 48 || ss[i] > 57 && ss[i] < 65 || ss[i] > 90 && ss[i] < 97 || ss[i] > 122 {
			return false
		}
	}
	return true
}
