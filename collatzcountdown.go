package piscine

func CollatzCountdown(n int) int {
	sd := 0
	if n <= 0 {
		return -1
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		sd++
	}
	return sd
}
