package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var co int
	for l != nil {
		if co == pos {
			return l
		}
		co++
		l = l.Next
	}
	return nil
}
