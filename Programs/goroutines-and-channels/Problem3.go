package goroutinesandchannels

// even odd without synchronization


import (
	"context"
	"fmt"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	const (
		nums = 15
	)

	evenCh := make(chan int)
	oddCh := make(chan int)

	go printEven(ctx, evenCh)
	go printOdd(ctx, oddCh)

	for i := 1; i <= nums; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("main context is closed")
		default:
			if i%2 != 0 {
				oddCh <- i
			} else {

				evenCh <- i
			}
		}

	}

	close(evenCh)
	close(oddCh)
	cancel()

}

func printEven(ctx context.Context, evenCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is cancelled")
			return
		case num, ok := <-evenCh:
			if !ok {
				fmt.Println("channel is closed")
				return
			}

			fmt.Println(num)
		}
	}
}

func printOdd(ctx context.Context, oddCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is cancelled")
			return
		case num, ok := <-oddCh:
			if !ok {
				fmt.Println("channel is closed")
				return
			}

			fmt.Println(num)
		}
	}
}
