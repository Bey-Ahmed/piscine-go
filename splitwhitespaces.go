package piscine

func SplitWhiteSpaces(s string) []string {
	size := len(s)

	if size <= 0 {
		return nil
	}

	split := []string{}
	start := 0
	end := 0
	spaces := 0
	from := 0
	for from < size && (s[from] == ' ' || s[from] == '\t' || s[from] == '\n') {
		from++
	}
	for i := from; i < size-1; i++ {
		if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n') && s[i+1] != ' ' && s[i+1] != '\t' && s[i+1] != '\n' {
			end = i - spaces
			spaces = 0
			split = append(split, s[start:end])
			start = i + 1
		} else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n') && (s[i+1] == ' ' || s[i+1] == '\t' || s[i+1] == '\n') {
			spaces++
		}
	}
	if s[size-1] == ' ' || s[size-1] == '\t' || s[size-1] == '\n' {
		end = size - 1 - spaces
		split = append(split, s[start:end])
	} else {
		split = append(split, s[end+1:size])
	}

	return split
}
