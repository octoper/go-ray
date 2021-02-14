package payloads

import (
	"strings"
)

func NewImagePayload(location string) Payload {
	location = strings.ReplaceAll(location, "\"", "")

	return NewCustomPayload("<img src=\""+ location +"\" alt=\"\" />", "Image")
}