package piscine

func Index(s, b string) int {
	ss := []rune(s)
	bb := []rune(b)
	var sss int
	var bbb int

	for range ss {
		sss++
	}
	for range bb {
		bbb++
	}
	for i := 0; i < sss-bbb; i++ {
		if b == s[i:i+bbb] {
			return i
		}
	}
	return -1
}
