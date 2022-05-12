package piscine

func Compact(ptr *[]string) int {
	sd := []string{}
	for _, c := range *ptr {
		if c != "" {
			sd = append(sd, c)
		}
	}
	*ptr = sd
	return len(sd)
}
