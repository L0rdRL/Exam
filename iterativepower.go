package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	sd := nb
	for i := 1; i < power; i++ {
		sd *= nb
	}
	return sd
}
