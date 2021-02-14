package payloads

import (
	"github.com/octoper/ray/utils"
	"path"
	"strings"
)

func NewImagePayload(location string) Payload {
	if utils.FileExists(location) {
		location = path.Join(path.Dir(location))
	}

	location = strings.ReplaceAll(location, "\"", "")

	return NewCustomPayload("<img src=\""+ location +"\" alt=\"\" />", "Image")
}