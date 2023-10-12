package handler

import (
	"fmt"
	"io"
	"net/http"
)

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	messageUrl := r.FormValue("messageUrl")

	response, err := http.Get(messageUrl)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed. Code:", response.StatusCode)
		return
	}

	message, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Problem parsing file:", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
