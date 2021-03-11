package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Client struct
type Client struct {
	host string
	port int
}

// LockRespnse struct
type LockRespnse struct {
	Name          string `json:"name"`
	Active        bool   `json:"active"`
	StopExecution bool   `json:"stop_execution"`
	GroupUuid     string `json:"displayed_on_group_uuid"`
}


// NewClient creates a Client instance
func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetPort(port int) int {
	c.port = port
	return c.port
}

func (c *Client) SetHost(host string) string {
	c.host = host
	return c.host
}


func (c *Client) Sent(requestPayload interface{}) (*http.Response, error) {
	requestJson, _ := json.Marshal(requestPayload)

	responseBody := bytes.NewBuffer(requestJson)

	//Make a request to Ray
	resp, err := http.Post("http://" + c.host + ":" + strconv.Itoa(c.port), "application/json", responseBody)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp, err
}

func (c *Client) LockExists(lockName string) LockRespnse {
	//Make a request to Ray
	resp, err := http.Get("http://" + c.host + ":" + strconv.Itoa(c.port) + "/locks/"+ lockName)

	if err != nil {
		panic(err)
	}

	body, bodyErr := io.ReadAll(resp.Body)

	if bodyErr != nil {
		panic(bodyErr)
	}

	res := LockRespnse{}
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		panic(err)
	}

	if res.StopExecution {
		os.Exit(1)
	}

	return res
}
