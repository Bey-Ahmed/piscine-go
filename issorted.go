package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	size := len(a)
	if size < 2 {
		return true
	}
	if f(a[0], a[size-1]) <= 0 {
		for i := 0; i < size-1; i++ {
			for j := i + 1; j < size; j++ {
				if f(a[i], a[j]) > 0 {
					return false
				}
			}
		}
	} else {
		for i := 0; i < size-1; i++ {
			for j := i + 1; j < size; j++ {
				if f(a[i], a[j]) < 0 {
					return false
				}
			}
		}
	}
	return true
}
