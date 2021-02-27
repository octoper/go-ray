package ray

// Green Color
func (r *application) Green() *application {
	return r.SendRequest(NewColorPayload("green"))
}

// Orange Color
func (r *application) Orange() *application {
	return r.SendRequest(NewColorPayload("orange"))
}

// Red Color
func (r *application) Red() *application {
	return r.SendRequest(NewColorPayload("red"))
}

// Blue Color
func (r *application) Blue() *application {
	return r.SendRequest(NewColorPayload("blue"))
}

// Purple Color
func (r *application) Purple() *application {
	return r.SendRequest(NewColorPayload("purple"))
}

// Gray Color
func (r *application) Gray() *application {
	return r.SendRequest(NewColorPayload("gray"))
}