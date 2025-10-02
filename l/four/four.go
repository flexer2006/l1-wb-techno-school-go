package four

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("sleep")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func deadline() {
	var wg sync.WaitGroup
	ctx, stp := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stp()
	for range 5 {
		wg.Go(func() {
			worker(ctx)
		})
	}
	<-ctx.Done()
	wg.Wait()
	fmt.Println(1_2_3)
}
