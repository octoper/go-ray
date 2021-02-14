package ray

import "github.com/octoper/go-ray/payloads"

func (r *application) Small() *application {
	return r.SendRequest(payloads.NewSizePayload("sm"))
}

func (r *application) Large() *application {
	return r.SendRequest(payloads.NewSizePayload("lg"))
}
