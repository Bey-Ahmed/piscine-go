package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	keep := start
	counter := 0
	for keep != 1 {
		if keep%2 == 0 {
			keep /= 2
		} else {
			keep = 3*keep + 1
		}
		counter++
	}
	return counter
}
