package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)


func emailWorker(ch chan Recipient,id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost:="localhost"
		smtpPort:="1025"

		// formatingMsg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s",
		// 	"vinit@gmail.com",
		// 	recipient.Email,
		// 	"testing")
		// msg:= []byte(formatingMsg)

		msg,err:=executeTemplate(recipient)
		if err != nil{
			log.Fatal(err)
		}
		err = smtp.SendMail(smtpHost+":"+smtpPort,nil,"vinit@gmail.com",[]string{recipient.Email},[]byte(msg)) 
		if err != nil{
			log.Fatal(err)
		}

		time.Sleep(50*time.Millisecond)
		fmt.Println(id,recipient)
	}
}