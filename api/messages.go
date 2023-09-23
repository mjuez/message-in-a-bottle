package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"regexp"
	"time"
)

type Message struct {
	Date     string `json:"date"`
	ImageUrl string `json:"image_url"`
	TextUrl  string `json:"text_url"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	url := "https://filedn.com/lQG3rKUjKEekfVlDSgDuyvR/message_in_a_bottle/messages/"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Can't get the URL", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed. Code:", response.StatusCode)
		return
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Cannot read HTTP response body", err)
		return
	}

	bodyText := string(bodyBytes)
	regex := `"name": "([^"]+\.(jpg|jpeg|png|gif))"`
	compiledRegex := regexp.MustCompile(regex)
	matches := compiledRegex.FindAllStringSubmatch(bodyText, -1)
	messages := []Message{}
	for _, match := range matches {
		if len(match) >= 2 {
			message := match[1]
			format := filepath.Ext(message)
			messageName := message[:len(message)-len(format)]

			date, err := time.Parse("060102", messageName)
			if err != nil {
				fmt.Println("Cannot analyse date:", err)
				return
			}

			jsonMessage := Message{
				Date:     date.Format("02/01/06"),
				ImageUrl: url + message,
				TextUrl:  url + messageName + ".txt",
			}
			messages = append(messages, jsonMessage)
		}
	}

	numMessages := len(messages)
	invertedMessages := make([]Message, numMessages)
	for i, message := range messages {
		invertedMessages[numMessages-1-i] = message
	}

	w.Header().Set("Content-Type", "application/json")
	encodingError := json.NewEncoder(w).Encode(invertedMessages)
	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
		return
	}
}
