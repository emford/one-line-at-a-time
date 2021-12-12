package main

import (
	"log"
	"os"
)

// link to subscribe
// make server
// tests
// cleanup docs
// lint
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("not enough arguments, need number to send to")
	}
	toNumber := os.Args[1]
	err := GetMessaging(toNumber)
	if err != nil {
		log.Fatalf("failed getting message %w", err)
	}
	message := "Task completed."
	log.Println(message)
}
