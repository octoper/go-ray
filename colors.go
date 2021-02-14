package ray

import "github.com/octoper/ray/payloads"

// Green Color
func (r *Application) Green() *Application {
	return r.SendRequest(payloads.NewColorPayload("green"))
}

// Orange Color
func (r *Application) Orange() *Application {
	return r.SendRequest(payloads.NewColorPayload("orange"))
}

// Red Color
func (r *Application) Red() *Application {
	return r.SendRequest(payloads.NewColorPayload("red"))
}

// Blue Color
func (r *Application) Blue() *Application {
	return r.SendRequest(payloads.NewColorPayload("blue"))
}

// Purple Color
func (r *Application) Purple() *Application {
	return r.SendRequest(payloads.NewColorPayload("purple"))
}

// Gray Color
func (r *Application) Gray() *Application {
	return r.SendRequest(payloads.NewColorPayload("gray"))
}