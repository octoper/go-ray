package ray

type Payload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
	Origin interface{} `json:"origin"`
}
