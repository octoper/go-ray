package payloads

import "encoding/json"

// NewJsonStringPayload creates a new Json String Payload
func NewJsonStringPayload(value interface{}) Payload {
	jsonValue, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	return Payload {
		Type: "json_string",
		Content: map[string]interface{} {
			"value": string(jsonValue),
		},
	}
}
