package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	courant := l
	for courant != nil {
		check := courant.Next
		for check != nil {
			if courant.Data > check.Data {
				courant.Data, check.Data = check.Data, courant.Data
			}
			check = check.Next
		}
		courant = courant.Next
	}

	return l
}
