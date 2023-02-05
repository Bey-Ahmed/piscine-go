package piscine

func IsNumeric(s string) bool {
	if len(s) <= 0 {
		return false
	}
	store := []rune(s)
	size := len(s)
	start := 0
	if store[0] == '-' {
		start = 1
	}
	for i := start; i < size; i++ {
		if store[i] < '0' || store[i] > '9' {
			return false
		}
	}
	return true
}
