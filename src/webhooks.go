package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	Messaging "github.com/hbollon/gobot/src/messaging"
	Yaml "github.com/hbollon/gobot/src/yaml"
)

var config Yaml.Config // Config struct instance
var configInit = Yaml.GetConfig()
var jsonData = json.Unmarshal([]byte(configInit), &config)

// Callback : callback struct for responses
// Based on Facebook API callback standards
type Callback struct {
	Object string `json:"object,omitempty"`
	Entry  []struct {
		ID        string                `json:"id,omitempty"`
		Time      int                   `json:"time,omitempty"`
		Messaging []Messaging.Messaging `json:"messaging,omitempty"`
	} `json:"entry,omitempty"`
}

// Handle GET request made to webhook
func whGetHandler(writter http.ResponseWriter, request *http.Request) {
	token := config.VerifyToken
	// Getting 'verify_token' and 'challenge' from incoming Messenger webhook request
	reqToken := request.URL.Query().Get("hub.verify_token")
	reqChallenge := request.URL.Query().Get("hub.challenge")

	// Check if req token match with config file token
	if reqToken == token {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusOK)
		writter.Write([]byte(reqChallenge))
	} else {
		fmt.Fprint(writter, "Tokens don't match")
	}
}

// Handle POST request made to webhook
// It serialize request content into Callback and call processMassage for response
func whPostHandler(writter http.ResponseWriter, request *http.Request) {
	fmt.Printf("POST request received\n")

	var callback Callback
	json.NewDecoder(request.Body).Decode(&callback) // Read and parse request body into callback

	// Check if req token match with config file token
	if callback.Object == "page" {
		for _, data := range callback.Entry {
			for _, ev := range data.Messaging {
				Messaging.MessageBuilder(ev, config.AccessToken, config.MatchPercentage)
			}
		}
		writter.WriteHeader(200)
		writter.Write([]byte("Message received !"))
	} else {
		writter.WriteHeader(404)
		writter.Write([]byte("Message not supported !"))
	}
}
