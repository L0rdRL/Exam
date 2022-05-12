package piscine

func AlphaCount(s string) int {
	ss := 0
	b := []rune(s)
	for i := 0; i < len(s); i++ {
		if b[i] >= 97 && b[i] <= 122 || b[i] >= 65 && b[i] <= 90 {
			ss += 1
		}
	}
	return ss
}
