package piscine

func ListRemoveIf(l *List, d interface{}) {
	prev := l.Head
	if l.Head != nil && l.Head.Data == d {
		l.Head = l.Head.Next
	}
	it := l.Head
	for it != nil {
		if it.Data != d {
			prev = it
		}
		prev.Next = it.Next
		it = it.Next
	}
}
