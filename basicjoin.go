package piscine

func BasicJoin(elems []string) string {
	if len(elems) == 0 {
		return ""
	}

	concat := elems[0]
	size := len(elems)
	for i := 1; i < size; i++ {
		concat += elems[i]
	}

	return concat
}
