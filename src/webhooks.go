package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	Messaging "gobot/src/messaging"
)

var config Config // Config struct instance
var tokenInit = getToken()
var jsonData = json.Unmarshal([]byte(tokenInit), &config)

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
/* func whPostHandler(writter http.ResponseWriter, request *http.Request) {
	var callback Callback

	// Check if req token match with config file token
	if reqToken == token {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusOK)
		writter.Write([]byte(reqChallenge))
	} else {
		fmt.Fprint(writter, "Tokens don't match")
	}
} */
