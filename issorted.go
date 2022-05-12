package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var sd []int
	for i := 0; i < len(a)-1; i++ {
		sd = append(sd, f(a[i], a[i+1]))
	}
	countpos, countneg := 0, 0
	for _, c := range sd {
		if c < 0 {
			countpos++
		} else if c > 0 {
			countneg++
		}
	}
	if countpos == 0 || countneg == 0 {
		return true
	}
	return false
}
