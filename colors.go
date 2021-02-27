package ray

import "github.com/octoper/go-ray/payloads"

// Green Color
func (r *application) Green() *application {
	return r.SendRequest(payloads.NewColorPayload("green"))
}

// Orange Color
func (r *application) Orange() *application {
	return r.SendRequest(payloads.NewColorPayload("orange"))
}

// Red Color
func (r *application) Red() *application {
	return r.SendRequest(payloads.NewColorPayload("red"))
}

// Blue Color
func (r *application) Blue() *application {
	return r.SendRequest(payloads.NewColorPayload("blue"))
}

// Purple Color
func (r *application) Purple() *application {
	return r.SendRequest(payloads.NewColorPayload("purple"))
}

// Gray Color
func (r *application) Gray() *application {
	return r.SendRequest(payloads.NewColorPayload("gray"))
}