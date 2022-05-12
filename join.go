package piscine

func Join(t []string, d string) string {
	var b string
	for i := 0; i < len(t)-1; i++ {
		b = b + t[i] + d
	}
	if len(t) > 1 {
		b = b + t[len(t)-1]
	}
	return b
}
