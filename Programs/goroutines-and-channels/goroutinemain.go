package goroutinesandchannels

func EnableGoroutine() {
	RunGoroutine()

	RunWorkers()
}

func RunGoroutine() {
	go func() {
		SayHello("Pratik Gaikwad")
	}()
}
