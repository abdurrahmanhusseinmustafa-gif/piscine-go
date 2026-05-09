package piscine

func ListSize(l *List) int {
	count := 0
	it := l.Head
	for it != nil {
		count++
		it = it.Next
	}
	return count
}
