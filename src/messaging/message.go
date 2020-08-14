package messaging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Facebook API url for request
const API_URL = "https://graph.facebook.com/v8.0/me/messages?access_token=%s"

// Build response request from Messaging
func MessageBuilder(ev Messaging) {
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: ev.Sender.ID,
		},
		Message: Message{
			Text: "Test response",
		},
	}

	// Building response request
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(&response)
	request, err := http.NewRequest("POST", fmt.Sprintf(API_URL, os.Getenv("PAGE_ACCESS_TOKEN")), requestBody)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Panic(err)
	}

	// Send final request by http
	res, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}

	defer res.Body.Close()
}
