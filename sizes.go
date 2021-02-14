package ray

import "github.com/octoper/go-ray/payloads"

func (r *Application) Small() *Application {
	return r.SendRequest(payloads.NewSizePayload("sm"))
}

func (r *Application) Large() *Application {
	return r.SendRequest(payloads.NewSizePayload("lg"))
}
