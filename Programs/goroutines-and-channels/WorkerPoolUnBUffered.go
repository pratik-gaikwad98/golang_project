package goroutinesandchannels

import (
	"fmt"
	"sync"
	"time"
)

func WorkerPoolUnBuffered() {
	var wg sync.WaitGroup

	const workers = 3
	const JobCount = 13
	job := make(chan int)
	result := make(chan int)

	//start workers
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go DoWorkerJobForUnbuffered(i, &wg, job, result)
	}

	//start producer
	go func() {
		for i := 1; i <= JobCount; i++ {
			job <- i
		}

		//All jobs are send now we close the channel no more jobs are comming
		close(job)

	}()

	//Now we will wait till all goroutine finshes
	go func() {
		wg.Wait()
		// once all go routine completed we will close result channell too
		close(result)
	}()

	//now we will loop over result to consume the data

	for res := range result {
		fmt.Println("result is ", res)
	}

}

func DoWorkerJobForUnbuffered(workerId int, wg *sync.WaitGroup, job <-chan int, result chan<- int) {
	defer wg.Done()

	for j := range job {
		time.Sleep(time.Second)
		fmt.Printf("worker %d is processing job %d\n", workerId, j)
		result <- j
	}
}
