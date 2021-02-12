package payloads

import "encoding/json"

type JsonStringPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewJsonStringPayload(value string) *JsonStringPayload {
	return &JsonStringPayload{
		Type: "json_string",
		Content: map[string]interface{} {
			"value": json.Marshal(value),
		},
	}
}