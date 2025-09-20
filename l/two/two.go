package two

import (
	"fmt"
	"sync"
)

func concurrencyQuadro() {
	arr := []int{2, 4, 6, 8, 10}
	res := make([]int, len(arr))

	var wg sync.WaitGroup

	for i := range arr {
		idx := i
		wg.Go(func() {
			res[idx] = arr[idx] * arr[idx]
		})
	}

	wg.Wait()

	for _, v := range res {
		fmt.Println(v)
	}
}
