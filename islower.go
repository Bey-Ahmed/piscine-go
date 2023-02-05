package piscine

func IsLower(s string) bool {
	if len(s) <= 0 {
		return false
	}
	store := []rune(s)
	size := len(s)
	for i := 0; i < size; i++ {
		if store[i] < 'a' || store[i] > 'z' {
			return false
		}
	}
	return true
}
