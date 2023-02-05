package piscine

func Split(s, sep string) []string {
	size := len(s)
	sepSize := len(sep)

	if size < sepSize || sepSize < 1 {
		return nil
	}

	intervals := 0
	for i := 0; i <= size-sepSize*2; i++ {
		if s[i:i+sepSize] == sep && s[i+sepSize:i+sepSize*2] != sep {
			intervals++
		}
	}

	split := make([]string, intervals+1)
	start := 0
	end := 0
	spaces := 0
	index := 0
	from := 0
	for from <= size-sepSize && s[from:from+sepSize] == sep {
		from++
	}
	for j := 0; j <= size-sepSize*2; j++ {
		if s[j:j+sepSize] == sep && s[j+sepSize:j+sepSize*2] != sep {
			end = j - spaces
			spaces = 0
			split[index] = s[start:end]
			index++
			start = j + sepSize
		} else if s[j:j+sepSize] == sep && s[j+sepSize:j+sepSize*2] == sep {
			spaces++
		}
	}
	if s[size-sepSize:size] == sep {
		end = size - sepSize - spaces
		split[index] = s[start:end]
	} else {
		split[index] = s[end+sepSize : size]
	}

	return split
}
