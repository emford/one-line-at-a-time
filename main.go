package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("not enough arguments, need number to send to")
	}
	toNumber := os.Args[1]
	err := GetMessaging(toNumber)
	if err != nil {
		log.Fatalf("failed getting message %s", err)
	}
	message := "Task completed."
	log.Println(message)
}
