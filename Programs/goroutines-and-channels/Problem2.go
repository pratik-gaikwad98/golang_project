package goroutinesandchannels

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*

Level 1 (Must solve in under 20 minutes)


Problem 1: Producer–Consumer

Create:

1 producer
3 consumers

Requirements:

Producer sends numbers 1–20.
Consumers print them.
Exit gracefully when the producer finishes.

Topics:

Channels
range
WaitGroup

*/

func workerP2(ctx context.Context, workerId int, job <-chan int, wg *sync.WaitGroup, mu *sync.Mutex, processedCount *int) {
	defer recoverF()
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("The worker ctx ix cancelled")
			return
		case worker, ok := <-job:

			if !ok {
				fmt.Println("Channel is closed")
				return
			}

			fmt.Printf("worker %d is processing %d job\n", workerId, worker)
			time.Sleep(time.Millisecond * 100)
			mu.Lock()
			*processedCount++
			mu.Unlock()

		}
	}

}

func ProblemMain2() {

	// declare const
	const (
		workerCount = 3
		jobCount    = 20
	)

	var (
		mu             sync.Mutex
		wg             sync.WaitGroup
		procesedNumber int
	)

	defer recoverF()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Millisecond*2000,
	)

	//declare channel
	job := make(chan int)

	// start worker

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go workerP2(ctx, i, job, &wg, &mu, &procesedNumber)
	}

	// send job to channel
	for i := 1; i <= jobCount; i++ {

		select {
		case <-ctx.Done():
			fmt.Println("main context is cancelled")
		case job <- i:
		}

	}

	// stop channel and wait here for goroutines
	close(job)
	wg.Wait()
	cancel()

	//print here
	fmt.Println("Processed number count is", procesedNumber)

}

func recoverF() {
	if r := recover(); r != nil {
		fmt.Println("recovered: ", r)
	}
}
