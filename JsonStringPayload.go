package ray

import "encoding/json"

func NewJsonStringPayload(value interface{}) Payload {
	jsonValue, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	return Payload{
		Type: "json_string",
		Content: map[string]interface{} {
			"value": string(jsonValue),
		},
	}
}