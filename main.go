package main

import (
	"log"
)

func main() {
	response, err := GetMessaging()
	if err != nil {
		log.Fatalf("failed getting message %w", err)
	}
	log.Printf("message response: %s", response)
	message := "Task completed."
	log.Println(message)
}
