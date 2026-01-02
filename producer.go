package main

import (
	"encoding/csv"
	"os"
	"fmt"
)

func loadRecipients(filePath string) error{
	f,err := os.Open(filePath)
	if err != nil {
		return err
	}
	r:= csv.NewReader(f)
	records,err = r.ReadAll()
	if err != nil {
		return err
	}
	for_,record:= range records[1:]{
		fmt.Println(record)
	}
}