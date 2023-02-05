package piscine

func AlphaCount(s string) int {
	count := 0
	size := len(s)
	for i := 0; i < size; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}
	return count
}
