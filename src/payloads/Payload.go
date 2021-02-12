package payloads

type Payload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}
