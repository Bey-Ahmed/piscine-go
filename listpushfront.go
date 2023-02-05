package piscine

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{
		Data: data,
		Next: nil,
	}

	if l.Head != nil {
		node.Next = l.Head
	}

	l.Head = node
}
