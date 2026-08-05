package goroutinesandchannels

import (
	"fmt"
	"sync"
	"time"
)

func DoWorkerJob(workerId int, wg *sync.WaitGroup, jobs <-chan int, result chan<- int) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d is processing %d job\n", workerId, job)
		time.Sleep(time.Second)
		result <- job * 2

	}

}

func WorkerPoolBuffered() {
	const WorkerCount = 3
	const Jobs = 19

	var wg sync.WaitGroup
	job := make(chan int, Jobs)
	result := make(chan int, Jobs)

	for i := 1; i <= WorkerCount; i++ {
		wg.Add(1)
		go DoWorkerJob(i, &wg, job, result)
	}

	for i := 1; i <= Jobs; i++ {
		job <- i
	}

	close(job)

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println("result", res)
	}

}
