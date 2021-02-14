package ray

var (
	rayHost = "127.0.0.1"
	rayPort = 23517
)

// Get the port Application is running
func (r *Application) Port() int {
	return r.port
}

// Set the port Application is running
func (r *Application) SetPort(port int) {
	rayPort = port
}

// Get the host Application is running
func (r *Application) Host() string {
	return r.host
}

// Set the host Application is running
func (r *Application) SetHost(host string) {
	rayHost = host
}
