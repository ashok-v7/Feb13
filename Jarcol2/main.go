// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Define the API URL at the package level
const apiURL = "https://api.example.com/data"

type ApiResponse struct {
	Message string `json:"message"`
}

// FetchMessage makes a GET request to the hardcoded URL and returns the message from the response.
func FetchMessage() (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResponse ApiResponse
	if err := json.Unmarshal(bodyBytes, &apiResponse); err != nil {
		return "", err
	}

	return apiResponse.Message, nil
}

func main() {
	message, err := FetchMessage()
	if err != nil {
		log.Fatalf("Error fetching message: %v", err)
	}
	fmt.Println("Fetched Message:", message)
}
