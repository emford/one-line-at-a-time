package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010" // import alias
)

// Client: Create a new Twilio client
func Client() (*twilio.RestClient, error) {
	client := twilio.NewRestClient()
	client.SetRegion("us1")
	client.SetEdge("umatilla")
	return client, nil
}

// GetFile: Open and read text file
func GetFile() ([]string, error) {
	file, err := os.Open("resume_text.txt")
	if err != nil {
		return nil, fmt.Errorf("failed setting up client: %w", err)

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text, nil
}

// GetMessaging: Create messaging params and return
// messages
func GetMessaging(toNumber string) error {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")

	client, err := Client()
	if err != nil {
		return fmt.Errorf("failed setting up client: %w", err) // string formatting? log.Log?
	}
	_ = client

	params := openapi.CreateMessageParams{}
	params.SetTo(toNumber)
	params.SetFrom(from)

	content, err := GetFile()
	if err != nil {
		return fmt.Errorf("failed to read text from file: %w", err)
	}
	for _, each_ln := range content {
		params.SetBody(each_ln)

		duration := time.Duration(len(each_ln) * 3)
		time.Sleep(duration * time.Second)

		resp, err := client.ApiV2010.CreateMessage(&params) // pointer reasons?
		if err != nil {
			return fmt.Errorf("failed creating message: %w", err)
		}
		price := "n/a"
		if resp.Price != nil {
			price = *resp.Price
		}
		log.Printf("sent message: cost of message = %s", price) // derefrencing?
	}
	return nil
}
