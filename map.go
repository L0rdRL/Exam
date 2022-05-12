package piscine

func Map(f func(int) bool, a []int) []bool {
	var s []bool
	for _, c := range a {
		s = append(s, f(c))
	}
	return s
}
