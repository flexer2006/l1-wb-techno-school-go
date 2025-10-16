package eleven

import "fmt"

func mnoj() {
	a, b, m := []int{1, 2, 3}, []int{2, 3, 4}, make(map[int]int)
	r := make([]int, 0)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		if m[v] > 0 {
			r = append(r, v)
			m[v]--
		}
	}

	fmt.Println(r)
}
