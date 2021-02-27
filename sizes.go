package ray

func (r *application) Small() *application {
	return r.SendRequest(NewSizePayload("sm"))
}

func (r *application) Large() *application {
	return r.SendRequest(NewSizePayload("lg"))
}
