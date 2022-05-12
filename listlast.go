package piscine

func ListLast(l *List) interface{} {
	if l != nil {
		if l.Tail != nil {
			return l.Tail.Data
		}
	}
	return nil
}
