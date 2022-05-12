package piscine

func SplitWhiteSpaces(r string) []string {
	s := r
	var sd []string
	d := 0
	for i := 0; i < len(s); i++ {
		if s[i] == byte(rune(32)) && s[i+1] == byte(rune(32)) {
			sd = append(sd, s[d:i])
			d = i + 2
			i = i + 1
		} else if i == len(s)-1 {
			sd = append(sd, s[d:])
		} else if s[i] == byte(rune(32)) {
			sd = append(sd, s[d:i])
			d = i + 1
			i = i + 1
		}
	}
	if rune(sd[0][0]) == rune(32) {
		return sd[1:]
	}

	return sd
}

/*SplitWhiteSpaces("F`W6m%BxHI{y/ UB Y'Rea%n9`o ~0H`L-8.Fa.KT ?StWT,^=J:WrG OWhXudsG]3H6' Fk a T~pw7}z$ $r]Kc=!W ,|K] {O\\'M::31K[LY")
[]string{"F`W6m%BxHI{y/", "UB", "Y'Rea%n9`o", "~0H`L-8.Fa.KT", "?StWT,^=J:WrG", "OWhXudsG]3H6'", "Fk", "a T~pw7}z$", "$r]Kc=!W", ",|K]", "{O\\'M::31K[LY"}
[]string{"F`W6m%BxHI{y/", "UB", "Y'Rea%n9`o", "~0H`L-8.Fa.KT", "?StWT,^=J:WrG", "OWhXudsG]3H6'", "Fk", "a", "T~pw7}z$", "$r]Kc=!W", ",|K]", "{O\\'M::31K[LY"}
*/
