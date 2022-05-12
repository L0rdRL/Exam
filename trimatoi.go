package piscine

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	sd := []rune(s)
	var numv int
	B := len(sd) - 1
	i := 0
	if sd[0] == 45 || sd[0] == 43 {
		i++
	}
	ss := 0

	for ; i <= B; i++ {
		if sd[i] == '-' && numv == 0 {
			ss++
		}

		if sd[i] < 48 || sd[i] > 57 {
			continue
		}
		numv *= 10
		numv += int(sd[i] - 48)

	}
	if ss > 0 {
		return -numv
	}
	return numv
}
