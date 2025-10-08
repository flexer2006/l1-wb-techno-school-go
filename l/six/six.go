package six

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func normalStopGoroutine() {
	num := 10

	go func() {
		if num%2 == 0 {
			return
		}
	}()

	go func() {
		fmt.Println(num)
	}()
}

func channelStopGoroutine() {
	done := make(chan struct{})
	go func() {
		<-done
	}()
	done <- struct{}{}

	ch := make(chan int)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	ch <- 1
	ch <- 2
	close(ch)

	bufCh := make(chan string, 1)
	go func() {
		for msg := range bufCh {
			if msg == "stop" {
				return
			}
			fmt.Println(msg)
		}
	}()
	bufCh <- "hello"
	bufCh <- "stop"
	close(bufCh)
}

func contextStopGoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
	}()
	cancel()

	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancelDeadline()
	go func() {
		<-ctxDeadline.Done()
	}()

	time.Sleep(2 * time.Second)
}

func timerTickerStopGoroutine() {
	var wg sync.WaitGroup

	wg.Go(func() {
		<-time.After(500 * time.Millisecond)
	})

	ticker := time.NewTicker(200 * time.Millisecond)
	wg.Go(func() {
		defer ticker.Stop()
		for range 3 {
			<-ticker.C
		}
	})

	timer := time.NewTimer(700 * time.Millisecond)
	stop := make(chan struct{})
	wg.Go(func() {
		for {
			select {
			case <-timer.C:
				return
			case <-stop:
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				return
			}
		}
	})
	close(stop)

	wg.Wait()
}

func runtimeStopGoroutine() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		runtime.Goexit()
	}()
	<-done
}

func panicStopGoroutine() {
	var wg sync.WaitGroup

	wg.Go(func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		panic("panic")
	})

	wg.Wait()
}

func syncPrimitiveStopGoroutine() {
	var wg sync.WaitGroup

	done := make(chan struct{})
	wg.Go(func() {
		<-done
	})
	close(done)

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	stop := false

	wg.Go(func() {
		mu.Lock()
		for !stop {
			cond.Wait()
		}
		mu.Unlock()
	})

	mu.Lock()
	stop = true
	mu.Unlock()
	cond.Broadcast()

	var m sync.Mutex
	stopped := false
	wg.Go(func() {
		for {
			m.Lock()
			s := stopped
			m.Unlock()
			if s {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	})

	m.Lock()
	stopped = true
	m.Unlock()

	wg.Wait()
}

func osSignalStopGoroutine() {
	var wg sync.WaitGroup

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	done := make(chan struct{})

	wg.Go(func() {
		<-done
	})

	go func() {
		<-sigCh
		close(done)
	}()

	_ = syscall.Kill(os.Getpid(), syscall.SIGINT)

	wg.Wait()
}
