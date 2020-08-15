package messaging

import (
	"bytes"
	"encoding/json"
	"fmt"
	Analysis "gobot/src/analysis"
	"log"
	"net/http"
)

// Facebook API url for request
const API_URL = "https://graph.facebook.com/v8.0/me/messages?access_token=%s"

// Build response request from Messaging
func MessageBuilder(ev Messaging, accessToken string) {
	fmt.Printf("Building response message...\n")
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: ev.Sender.ID,
		},
		Message: Message{
			Text: Analysis.FindResponse(ev.Message.Text),
		},
	}

	// Building response request
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(&response)
	request, err := http.NewRequest("POST", fmt.Sprintf(API_URL, accessToken), requestBody)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("Error during request creation !\n")
		log.Panic(err)
	}

	// Send final request by http
	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error during request sending !\n")
		log.Panic(err)
	}

	fmt.Printf("Done !\n")
	defer res.Body.Close()
}
