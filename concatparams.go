package piscine

func ConcatParams(args []string) string {
	size := len(args)

	concat := ""
	for i := 0; i < size; i++ {
		if i < size-1 {
			concat += args[i] + "\n"
		} else {
			concat += args[i]
		}
	}

	return concat
}
