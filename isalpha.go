package piscine

func IsAlpha(s string) bool {
	if len(s) <= 0 {
		return false
	}
	store := []rune(s)
	size := len(s)
	for i := 0; i < size; i++ {
		if !((store[i] >= 'a' && store[i] <= 'z') || (store[i] >= 'A' && store[i] <= 'Z') || (store[i] >= '0' && store[i] <= '9')) {
			return false
		}
	}
	return true
}
