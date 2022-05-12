package piscine

func Rot14(s string) string {
	runi := []rune(s)
	for i := 0; i < len(s); i++ {
		if runi[i] >= 65 && runi[i] <= 90 {
			if runi[i]+14 > 90 {
				runi[i] = runi[i] + 14 - 26
			} else {
				runi[i] = runi[i] + 14
			}
		}
		if runi[i] >= 97 && runi[i] <= 122 {
			if runi[i]+14 > 122 {
				runi[i] = runi[i] + 14 - 26
			} else {
				runi[i] = runi[i] + 14
			}
		}
	}
	return string(runi)
}
