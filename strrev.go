package piscine

func StrRev(s string) string {
	runeArray := []byte(s)
	reverseStr := []byte{}
	size := len(runeArray)

	for i := 0; i < size; i++ {
		reverseStr = append(reverseStr, runeArray[size-1-i])
	}

	return string(reverseStr)
}
