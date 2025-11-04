package twentyfree

func RemoveAt[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return nil
	}
	copy(s[i:], s[i+1:])
	var zero T
	s[len(s)-1] = zero
	s = s[:len(s)-1]
	return s[:len(s):len(s)]
}
