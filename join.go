package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	concat := strs[0]
	size := len(strs)
	for i := 1; i < size; i++ {
		concat += sep + strs[i]
	}

	return concat
}
