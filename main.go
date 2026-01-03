package main

import "sync"

type Recipient struct {
	Name string
	Email string
}

func main() {
	recipientsChannel := make(chan Recipient)

	go loadRecipients("recipients.csv", recipientsChannel)

	var wg sync.WaitGroup
	workerCount := 10
	for i := 1; i < workerCount; i++ {
		wg.Add(1)
		go emailWorker(recipientsChannel,i,&wg)
	}

	wg.Wait()
}