package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return &NodeI{
			Data: data_ref,
			Next: nil,
		}
	}

	if l.Data > data_ref {
		node := &NodeI{
			Data: data_ref,
			Next: l,
		}

		l = node

		return l
	}

	courant := l
	for courant.Next != nil && courant.Next.Data <= data_ref {
		courant = courant.Next
	}

	node := &NodeI{
		Data: data_ref,
		Next: nil,
	}

	if courant.Next != nil {
		node.Next = courant.Next
	}

	courant.Next = node

	return l
}
