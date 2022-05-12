package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	s := make([]int, max-min)
	j := min

	for i := 0; i < len(s); i++ {
		s[i] = j + i
	}
	return s
}
