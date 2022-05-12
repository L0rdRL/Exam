package piscine

func BasicJoin(s []string) string {
	var a string
	for i := 0; i < len(s); i++ {
		a = a + s[i]
	}
	return a
}
