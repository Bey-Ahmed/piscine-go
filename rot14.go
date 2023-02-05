package piscine

func Rot14(s string) string {
	size := len(s)

	if size <= 0 {
		return ""
	}

	rot := make([]rune, size)
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			if (char+14)/'z' == 1 && (char+14)%'z' == 0 {
				rot[i] = char + 14
			} else {
				rot[i] = ('a'-1)*((char+14)/'z') + (char+14)%'z'
			}
		} else if char >= 'A' && char <= 'Z' {
			if (char+14)/'Z' == 1 && (char+14)%'Z' == 0 {
				rot[i] = char + 14
			} else {
				rot[i] = ('A'-1)*((char+14)/'Z') + (char+14)%'Z'
			}
		} else {
			rot[i] = char
		}
	}

	return string(rot)
}
