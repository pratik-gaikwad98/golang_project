package goroutinesandchannels

// even odd with synchro

import (
	"context"
	"fmt"
	"sync"
)

func ProblemMain4() {
	ctx, _ := context.WithCancel(context.Background())
	var (
		nums = 20
		wg   sync.WaitGroup
	)

	evenCh := make(chan int)
	oddCh := make(chan int)
	wg.Add(2)
	go printEvenNum(ctx, &wg, evenCh, oddCh, nums)
	go printOddNum(ctx, &wg, oddCh, evenCh, nums)

	//sending 1st number to odd
	evenCh <- 1

	wg.Wait()

}

func printEvenNum(ctx context.Context, wg *sync.WaitGroup, evenCh <-chan int, oddCh chan<- int, nums int) {
	defer wg.Done()

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

			if num == nums {
				fmt.Println("stoping on", num)
				close(oddCh)
				return
			}
			fmt.Println(num)
			num++
			oddCh <- num
		}
	}
}

func printOddNum(ctx context.Context, wg *sync.WaitGroup, oddCh <-chan int, evenCh chan<- int, nums int) {

	defer wg.Done()
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

			if num == nums {
				fmt.Println("stoping on", num)
				close(evenCh)
				return
			}

			fmt.Println(num)
			num++
			evenCh <- num
		}
	}
}
