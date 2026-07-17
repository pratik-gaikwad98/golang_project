package goroutinesandchannels

/*

Create a jobs channel.
	Start 3 worker goroutines.
	Each worker should:
	receive a job from the channel,
	print "Worker X processing Job Y"
	sleep for 1 second,
	continue until the channel is closed.
	Send 10 jobs into the channel.
	Use a sync.WaitGroup so main waits for all workers to finish instead of using time.Sleep().

*/

func RunWorkerAndJobs(){



}

