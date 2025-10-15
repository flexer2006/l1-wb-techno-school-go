package nine

import "fmt"

func generation(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, v := range nums {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func doubles(in <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for v := range in {
			ch <- v * 2
		}
		close(ch)
	}()

	return ch
}

func conveer() {
	for v := range doubles(generation(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)) {
		fmt.Println(v)
	}
}
