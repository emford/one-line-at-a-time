package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010" // import alias
)

// Create a new client
// if more than one thing is returned, need parenth.
func Client() (*twilio.RestClient, error) {
	client := twilio.NewRestClient()
	client.SetRegion("us1")
	client.SetEdge("umatilla")
	return client, nil
}

// Is preffered way to bubble up
func GetMessaging() ([]byte, error) {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")
	// subaccountSid := os.Getenv("TWILIO_SUBACCOUNT_SID")

	client, err := Client()
	if err != nil {
		return nil, fmt.Errorf("failed setting up client: %w", err) // string formatting? log.Log?
	}

	params := openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := client.ApiV2010.CreateMessage(&params) // pointer reasons?
	if err != nil {
		return nil, fmt.Errorf("failed creating message: %w", err)
	}
	response, err := json.Marshal(*resp) // always error on sad path - no use of else
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}
	return response, nil
}
