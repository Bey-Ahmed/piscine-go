package piscine

func SortWordArr(a []string) {
	size := len(a)

	if size <= 0 {
		return
	}

	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
