package piscine

func StrRev(s string) string {
	str := s
	bstr := []rune(str)
	for i, j := 0, len(bstr)-1; i < j; i, j = i+1, j-1 {
		bstr[i], bstr[j] = bstr[j], bstr[i]
	}
	return string(bstr)
}
