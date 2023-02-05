package piscine

func IsPrintable(s string) bool {
	if len(s) <= 0 {
		return false
	}

	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] < ' ' {
			return false
		}
	}
	return true
}
