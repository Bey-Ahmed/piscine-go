package piscine

func IsUpper(s string) bool {
	if len(s) <= 0 {
		return false
	}
	store := []rune(s)
	size := len(s)
	for i := 0; i < size; i++ {
		if store[i] < 'A' || store[i] > 'Z' {
			return false
		}
	}
	return true
}
