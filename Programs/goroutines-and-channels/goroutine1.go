package goroutinesandchannels

import (
	"fmt"
	"time"
)

func RunWorkers() {
	fmt.Println("Running workers")
	ch := make(chan string)
	workers := []string{"worker1", "worker2", "worker3"}

	// 1. Start a goroutine that strictly sends data to the channel
	go func() {
		for _, wrk := range workers {
			ch <- wrk
			// SayHello(<-ch)
		}
		close(ch)
	}()

	// 2. Consume from the channel on the main thread
	for wrk := range ch {
		SayHello(wrk) // Pass the string value instead of the channel
	}

}

func SayHello(ch string) {
	fmt.Printf("Hello from %s\n", ch)
	time.Sleep(1 * time.Second)
}
