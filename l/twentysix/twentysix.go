package twentysix

import "unicode"

func HasAllUniqueRunes(s string) bool {
	seen := make(map[rune]struct{}, len(s))
	for _, r := range s {
		r = unicode.ToLower(r)
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}
	return true
}
