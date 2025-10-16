package twelve

import (
	"fmt"
	"slices"
)

func uniqeSetMap() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, v := range animals {
		if _, exists := set[v]; exists {
			continue
		}
		set[v] = struct{}{}
	}

	for k := range set {
		fmt.Println(k)
	}
}

func uniqeSetSlice() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make([]string, 0)

	for _, v := range animals {
		f := slices.Contains(set, v)

		if !f {
			set = append(set, v)
		}
	}

	for _, v := range set {
		fmt.Println(v)
	}
}
