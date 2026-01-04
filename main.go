package main

import (
	"bytes"
	"sync"
	"text/template"
)

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


func executeTemplate(r Recipient) (string, error) {
	t,err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	t.Execute(&tpl,r)

	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}