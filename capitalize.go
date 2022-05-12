package piscine

func Capitalize(s string) string {
	ss := []rune(s)
	f := true
	for i := 0; i < len(s); i++ {
		if alnumeric((ss[i])) && f {
			if ss[i] >= 97 && ss[i] <= 122 {
				ss[i] = rune(ss[i] - 32)
			}
			f = false

		} else if ss[i] >= 65 && ss[i] <= 90 {
			ss[i] = rune(ss[i] + 32)
			f = false
		} else if !alnumeric(ss[i]) {
			f = true
		}
	}
	return string(ss)
}

func alnumeric(s rune) bool {
	if (s >= 65 && s <= 90) || (s >= 48 && s <= 57) || (s >= 97 && s <= 122) {
		return true
	}
	return false
}
