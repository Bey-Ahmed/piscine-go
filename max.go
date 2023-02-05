package piscine

func Max(a []int) int {
	size := len(a)

	if size <= 0 {
		return 0
	}

	max := a[0]
	for i := 1; i < size; i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}
