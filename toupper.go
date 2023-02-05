package piscine

func ToUpper(s string) string {
	store := []byte(s)
	size := len(s)
	for i := 0; i < size; i++ {
		if store[i] >= 'a' && store[i] <= 'z' {
			store[i] -= 'a' - 'A'
		}
	}
	return string(store)
}
