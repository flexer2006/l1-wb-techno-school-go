package five

import (
	"fmt"
	"time"
)

func timeoutTask() {
	ch := make(chan any, 1)
	go func() {
		defer close(ch)
		timeout := time.After(time.Duration(1 * time.Second))
		for {
			select {
			case <-timeout:
				return
			default:
				ch <- 1
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
