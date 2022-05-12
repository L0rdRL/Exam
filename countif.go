package piscine

func CountIf(f func(string) bool, tab []string) int {
	s := 0
	for _, c := range tab {
		if f(c) {
			s++
		}
	}
	return s
}
