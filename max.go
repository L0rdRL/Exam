package piscine

func Max(a []int) int {
	s := 0
	for _, c := range a {
		if s < c {
			s = c
		}
	}
	return s
}
