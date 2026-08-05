package goroutinesandchannels

/*

	5 worker goroutines.
	50 URLs (just use strings like "url1", "url2").
	Each worker:
	Reads one URL.
	Sleeps 500ms (simulate download).
	Randomly fails 20% of the time.
	Sends success/failure to a results channel.
	Main should:
	Count successful downloads.
	Count failed downloads.
	Print the final summary.
Use:
	WaitGroup
	Channels
	Worker Pool
	Context (cancel after 10 seconds)

*/

func RunWorkerAndJobs(){



}

