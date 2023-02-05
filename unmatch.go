package piscine

func Unmatch(a []int) int {
	size := len(a)

	if size <= 0 {
		return 0
	}

	if size == 1 {
		return a[0]
	}

	SortIntegerTable(a)

	for i := 0; i < size-1; i = i + 2 {
		if a[i] != a[i+1] {
			return a[i]
		}
	}

	if size%2 == 0 {
		return -1
	}

	return a[size-1]
}
