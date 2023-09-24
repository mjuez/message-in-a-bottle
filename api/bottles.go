package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type Bottle struct {
	Date     string `json:"date"`
	ImageUrl string `json:"image_url"`
	Message  string `json:"message"`
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
	regex := `"name": "((([^"\_]+)(\_\d+)?)\.(jpg|jpeg|png|gif))"`
	compiledRegex := regexp.MustCompile(regex)
	matches := compiledRegex.FindAllStringSubmatch(bodyText, -1)
	bottles := []Bottle{}
	for _, match := range matches {
		if len(match) >= 4 {
			bottle := match[1]
			bottleNoExt := match[2]
			date := match[3]

			parsedDate, err := time.Parse("060102", date)
			if err != nil {
				fmt.Println("Cannot analyse date:", err)
			} else {
				jsonBottle := Bottle{
					Date:     parsedDate.Format("02/01/06"),
					ImageUrl: url + bottle,
					Message:  openBottle(url + bottleNoExt + ".txt"),
				}
				bottles = append(bottles, jsonBottle)
			}
		}
	}

	numBottles := len(bottles)
	invertedBottles := make([]Bottle, numBottles)
	for i, bottle := range bottles {
		invertedBottles[numBottles-1-i] = bottle
	}

	w.Header().Set("Content-Type", "application/json")
	encodingError := json.NewEncoder(w).Encode(invertedBottles)
	if encodingError != nil {
		http.Error(w, encodingError.Error(), http.StatusInternalServerError)
		return
	}
}

func openBottle(messageUrl string) string {
	response, err := http.Get(messageUrl)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return ""
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed. Code:", response.StatusCode)
		return ""
	}

	message, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Problem parsing file:", err)
		return ""
	}

	return string(message)
}
