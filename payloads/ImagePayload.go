package payloads

import (
	"ray/utils"
	"strings"
)

func NewImagePayload(location string) *CustomPayload {
	if utils.FileExists(location) {
		location = "file://" + location
	}

	location = strings.ReplaceAll(location, "\"", "")

	return NewCustomPayload("<img src=\""+ location +"\" alt=\"\" />", "Image")
}