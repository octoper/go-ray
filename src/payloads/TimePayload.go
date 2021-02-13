package payloads

import "time"

func NewTimePayload(time time.Time, format string) Payload {
	timezoneName, _ := time.Zone()

	return Payload {
		Type: "carbon",
		Content: map[string]interface{} {
			"formatted": time.Format(format),
			"timestamp": time.Unix(),
			"timezone": timezoneName,
		},
	}
}