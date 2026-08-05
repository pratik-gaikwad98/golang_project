package goroutinesandchannels

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
You have to process 100 orders.

Requirements:

Spawn only 5 workers.
Workers should process orders concurrently.
Each order takes 100ms to process.
If processing takes more than 2 seconds, stop all workers immediately using context.
Print:
*/
type Order struct {
	ID int
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan Order,
	wg *sync.WaitGroup,
	processed *int,
	mu *sync.Mutex,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker ctx is cancelled")
			return
		case order, ok := <-jobs:
			if !ok {
				fmt.Println("channel is closed")
				return
			}
			fmt.Printf("Worker %d processing Order %d\n", id, order.ID)

			select {
			case <-ctx.Done():
				fmt.Println("channel is closed while processing order")
				return
			case <-time.After(time.Millisecond * 100):
				mu.Lock()
				*processed++
				mu.Unlock()
			}

		}
	}

}

func Problem1Main() {
	const (
		workerCount int = 5
		orderCount      = 100
	)
	var (
		wg             sync.WaitGroup
		mu             sync.Mutex
		processedCount int = 0
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Print("Recovered", r)
		}
	}()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second*2,
	)
	defer cancel()

	// Create jobs channel
	jobs := make(chan Order, workerCount)

	// Start 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg, &processedCount, &mu)
	}

	// Send 100 orders
	for i := 1; i <= orderCount; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("main context is cancelled")
			close(jobs)
			wg.Wait()
			return
		case jobs <- Order{i}:
		}
	}

	// Close channel
	close(jobs)
	wg.Wait()

	// Print processed count
	fmt.Println("Processed orders: ", processedCount)
}
