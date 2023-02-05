package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	size := len(a)

	if size <= 0 {
		return
	}

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
