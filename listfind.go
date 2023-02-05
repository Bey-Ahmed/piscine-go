package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	courant := l.Head
	for courant != nil && !comp(courant.Data, ref) {
		courant = courant.Next
	}
	if courant == nil {
		return nil
	}
	return &courant.Data
}
