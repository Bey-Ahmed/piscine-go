package piscine

func Index(s string, toFind string) int {
	if len(toFind) <= 0 {
		return 0
	}

	size := len(s)
	sizeToFind := len(toFind)
	for i := 0; i < size; i++ {
		if sizeToFind+i <= size-1 && s[i:i+sizeToFind] == toFind {
			return i
		}
	}
	return -1
}
