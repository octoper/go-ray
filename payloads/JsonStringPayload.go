package payloads

import "encoding/json"

func NewJsonStringPayload(value string) Payload {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return Payload{
		Type: "json_string",
		Content: map[string]interface{} {
			"value": string(json),
		},
	}
}