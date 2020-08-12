package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var config Config // Config struct instance
var tokenInit = getToken()
var jsonData = json.Unmarshal([]byte(tokenInit), &config)

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
