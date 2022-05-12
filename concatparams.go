package piscine

func ConcatParams(args []string) string {
	ss := args
	s := ""
	for i := 0; i < len(ss)-1; i++ {
		s += ss[i]
		s += string('\n')
	}
	if len(ss) > 1 {
		s += ss[len(ss)-1]
	}
	return s
}
