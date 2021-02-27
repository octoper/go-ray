package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type client struct {
	url string
}

// Create Client instance
func NewClient(url string) *client {
	newClient := client{
		url: url,
	}
	return &newClient
}


func (c *client) Sent(requestPayload interface{}) (*http.Response, error) {
	requestJson, _ := json.Marshal(requestPayload)

	//fmt.Println(string(requestJson))

	responseBody := bytes.NewBuffer(requestJson)

	//Make a request to Ray
	resp, err := http.Post(c.url, "application/json", responseBody)

	defer resp.Body.Close()

	return resp, err
}