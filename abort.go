package piscine

func Abort(a, b, c, d, e int) int {
	sd := []int{a, b, c, d, e}
	SortIntegerTable(sd)
	return sd[2]
}
