package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("invalid arguments, need number to send to")
		log.Fatal("usage: one-line-at-a-time <recipient-phone-number>")
	}
	toNumber := os.Args[1]
	err := GetMessaging(toNumber)
	if err != nil {
		log.Fatalf("failed getting message %s", err)
	}
	log.Println("Task completed.")
}
