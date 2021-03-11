package payloads

import (
	"strings"
)

// Create a new Image Payload
func NewImagePayload(location string) Payload {
	location = strings.ReplaceAll(location, "\"", "")

	return NewCustomPayload("<img src=\""+ location +"\" alt=\"\" />", "Image")
}
