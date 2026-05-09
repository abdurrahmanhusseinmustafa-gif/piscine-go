package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	current := l
	for i := 0; i < pos; i++ {
		if current == nil {
			return nil
		}
		current = current.Next
	}
	return current
}
