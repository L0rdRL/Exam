package piscine

func ListSize(l *List) int {
	var i int
	for l.Head != nil {
		if l.Head == nil {
			break
		} else {
			l.Head = l.Head.Next
			i++
		}
	}
	return i
}
