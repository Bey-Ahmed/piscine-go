package piscine

func Compact(ptr *[]string) int {
	size := len(*ptr)

	if size <= 0 {
		return 0
	}

	compact := []string{}
	for i := 0; i < size; i++ {
		if (*ptr)[i] != "" {
			compact = append(compact, (*ptr)[i])
		}
	}
	(*ptr) = compact
	return len(*ptr)
}
