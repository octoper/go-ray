package payloads

import "encoding/json"

type JsonStringPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewJsonStringPayload(value string) *JsonStringPayload {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return &JsonStringPayload{
		Type: "json_string",
		Content: map[string]interface{} {
			"value": string(json),
		},
	}
}