package eighteen

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	value atomic.Int64
}

func (c *SafeCounter) Inc() {
	c.value.Add(1)
}

func safeTask() {
	const workers = 10

	var (
		counter SafeCounter
		wg      sync.WaitGroup
	)

	for range workers {
		wg.Go(func() {
			counter.Inc()
			fmt.Println(counter.value.Load())
		})
	}

	wg.Wait()
	fmt.Println(counter.value.Load())
}
