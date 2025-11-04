package twenty

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var sb strings.Builder
	for i := len(strings.Fields(s)) - 1; i >= 0; i-- {
		sb.WriteString(strings.Fields(s)[i])
		sb.WriteByte(' ')
	}
	return sb.String()
}

func reverseWords2(s string) string {
	var (
		sb    strings.Builder
		first = true
	)

	sb.Grow(len(s))

	for i := len(s) - 1; i >= 0; {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		if i < 0 {
			break
		}
		end := i

		for i >= 0 && s[i] != ' ' {
			i--
		}
		start := i + 1

		if !first {
			sb.WriteByte(' ')
		}
		sb.WriteString(s[start : end+1])
		first = false
	}

	return sb.String()
}
func twentyTask() {
	fmt.Println(reverseWords("snow dog sun"))
}
