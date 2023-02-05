package piscine

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	store := []rune(s)
	size := len(s)

	if store[0] >= 'a' && store[0] <= 'z' {
		store[0] -= 'a' - 'A'
	}

	for i := 1; i < size; i++ {
		if !((store[i-1] >= 'a' && store[i-1] <= 'z') || (store[i-1] >= 'A' && store[i-1] <= 'Z') || (store[i-1] >= '0' && store[i-1] <= '9')) {
			if store[i] >= 'a' && store[i] <= 'z' {
				store[i] -= 'a' - 'A'
			}
		} else if store[i] >= 'A' && store[i] <= 'Z' {
			store[i] += 'a' - 'A'
		}
	}

	return string(store)
}
