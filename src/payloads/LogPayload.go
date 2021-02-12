package payloads

type LogPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewLogPayload(values interface{}) *LogPayload {
	return &LogPayload{
		Type: "log",
		Content: map[string]interface{} {
			"value": values,
		},
	}
}