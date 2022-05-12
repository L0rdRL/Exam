package piscine

func ListClear(l *List) {
	for l.Head != nil {
		l.Head = nil
		l.Tail = nil
	}
}
