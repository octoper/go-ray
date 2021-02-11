package payloads

type CreateLockPayload struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func NewCreateLockPayload(name string) *CreateLockPayload {
	return &CreateLockPayload{
		Type: "create_lock",
		Content: map[string]string {
			"name": name,
		},
	}
}