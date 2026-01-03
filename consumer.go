package main

import (
	"fmt"
	"sync"
)


func emailWorker(ch chan Recipient,id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		fmt.Println(id,recipient)
	}
}