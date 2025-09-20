package three

import (
	"fmt"
	"sync"
)

func worker(job <-chan int) {
	for j := range job {
		fmt.Printf("job = %d\n", j)
	}
}

func tasker() {
	workers := 5
	if workers < 1 {
		panic("workers < 1")
	}
	job := make(chan int, workers)
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			worker(job)
		})
	}
	for i := range workers {
		job <- i
	}
	close(job)
	wg.Wait()
}
