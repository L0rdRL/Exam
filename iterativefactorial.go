package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	fac := 1
	for i := nb; i > 0; i-- {
		fac = fac * i
	}
	return fac
}
