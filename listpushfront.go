package piscine

func ListPushFront(l *List, data interface{}) {
	second := l.Head
	l.Head = &NodeL{Data: data}
	l.Head.Next = second
}
