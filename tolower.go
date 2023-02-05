package piscine

func ToLower(s string) string {
	store := []byte(s)
	size := len(s)
	for i := 0; i < size; i++ {
		if store[i] >= 'A' && store[i] <= 'Z' {
			store[i] += 'a' - 'A'
		}
	}
	return string(store)
}
