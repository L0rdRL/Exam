package piscine

func Unmatch(a []int) int {
	for _, c := range a {
		con := 0

		for _, v := range a {
			if c == v {
				con++
			}
		}
		if con%2 == 1 {
			return c
		}
	}
	return -1
}
