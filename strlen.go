package piscine

func StrLen(s string) int {
	runeArray := []rune(s)
	size := len(runeArray)
	return size
}
