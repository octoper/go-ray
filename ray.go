package ray

import (
	"encoding/json"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

type Callable = func() bool

type application struct {
	uuid string
	host string
	port int
	enabled bool
	sentPayloads []Payload
}

type request struct {
	Uuid     string      `json:"uuid"`
	Payloads interface{} `json:"payloads"`
	Meta     map[string]string     `json:"meta"`
}

var applicationConfig = application{
	host: "127.0.0.1",
	port: 23517,
	enabled: true,
}

// Create New Ray instance
func NewRay() *application {
	ray := applicationConfig
	return &ray
}

// Get the UUID
func (r *application) Uuid() string {
	return r.uuid
}

// Get the port application is running
func (r *application) Port() int {
	return r.port
}

// Set the port application is running
func (r *application) SetPort(port int) {
	applicationConfig.port = port
}

// Get the host application is running
func (r *application) Host() string {
	return r.host
}

// Set the host application is running
func (r *application) SetHost(host string) {
	applicationConfig.host = host
}

// Get Snet Payloads as Json
func (r *application) SentJsonPayloads() ([]byte, error) {
	return json.Marshal(applicationConfig.sentPayloads)
}


func Ray(values ...interface{}) *application {
	r := NewRay()

	r.uuid = uuid.New().String()

	if values != nil {
		r.Send(values...)
	}

	return r
}

// Send Values
func (r *application) Send(values ...interface{}) *application {
	var payloadsMap []Payload

	for _, payload := range values {
		switch payload.(type) {
		case bool:
			payloadsMap = append(payloadsMap, NewBoolPayload(payload.(bool)))
		case nil:
			payloadsMap = append(payloadsMap, NewNullPayload())
		default:
			payloadsMap = append(payloadsMap, NewLogPayload(payload))
		}
	}

	return r.SendRequest(payloadsMap...)
}

// Enable Ray
func (r *application) Enable() {
	applicationConfig.enabled = true
}

// Disable Ray
func (r *application) Disable() {
	applicationConfig.enabled = false
}

// Create New Screen
func (r *application) NewScreen(name string) *application {
	return r.SendRequest(NewNewScreenPayload(name))
}

// Color
func (r *application) Color(color string) *application {
	return r.SendRequest(NewColorPayload(color))
}

// Send custom payload
func (r *application) SendCustom(content interface{}, label string) *application {
	return r.SendRequest(NewCustomPayload(content, label))
}

// Size
func (r *application) Size(size string) *application {
	return r.SendRequest(NewSizePayload(size))
}

// Hide
func (r *application) Hide() *application {
	return r.SendRequest(NewHidePayload())
}

// Hide App
func (r *application) HideApp() *application {
	return r.SendRequest(NewHideAppPayload())
}

// Show App
func (r *application) ShowApp() *application {
	return r.SendRequest(NewShowAppPayload())
}

// Clear Screen
func (r *application) ClearScreen() *application {
	return r.SendRequest(NewClearScreenPayload())
}

// Clear All
func (r *application) ClearAll() *application {
	return r.SendRequest(NewClearAllPayload())
}

// HTML
func (r *application) Html(html string) *application {
	return r.SendRequest(NewHtmlPayload(html))
}

// Notify
func (r *application) Notify(text string) *application {
	return r.SendRequest(NewNotifyPayload(text))
}

// Pass
func (r *application) Pass(arg interface{}) interface{} {
	r.Send(arg)
	return arg
}

// Boolean
func (r *application) Bool(bool bool) *application {
	return r.SendRequest(NewBoolPayload(bool))
}

// Null
func (r *application) Null() *application {
	return r.SendRequest(NewNullPayload())
}

// Charles
func (r *application) Charles() *application {
	return r.Send("🎶 🎹 🎷 🕺")
}

// String
func (r *application) String(str string) *application {
	return r.SendRequest(NewStringPayload(str))
}

//Time
func (r *application) Time(time time.Time) *application {
	return r.SendRequest(NewTimePayload(time, "2006-01-02 15:04:05"))
}

// Time with Format
func (r *application) TimeWithFormat(time time.Time, format string) *application {
	return r.SendRequest(NewTimePayload(time, format))
}

/**
 * Sends the provided value(s) encoded as a JSON string using json.Marshal.
 */
func (r *application) ToJson(jsons ...interface{}) *application {
	for _, jsonValue := range jsons {
		r.SendRequest(NewJsonStringPayload(jsonValue))
	}
	return r
}

// Image
func (r *application) Image(value string) *application {
	return r.SendRequest(NewImagePayload(value))
}

// Ban
func (r *application) Ban() *application {
	return r.Send("🕶")
}

// Die
func (r *application) Die() {
	os.Exit(1)
}

// Remove
func (r *application) Remove() *application {
	return r.SendRequest(NewRemovePayload())
}

// Show When
func (r *application) ShowWhen(show interface{}) *application {
	switch show.(type) {
	case Callable:
		show = show.(Callable)()
	}

	if !show.(bool) {
		return r.Remove()
	}

	return r
}

// Show If
func (r *application) ShowIf(show interface{}) *application {
	return r.ShowWhen(show)
}

// Remove When
func (r *application) RemoveWhen(show interface{}) *application {
	switch show.(type) {
	case Callable:
		show = show.(Callable)()
	}

	if show.(bool) {
		return r.Remove()
	}

	return r
}

// Remove If
func (r *application) RemoveIf(show interface{}) *application {
	return r.RemoveWhen(show)
}

// Set the host application is running
func (r *application) SendRequest(ResponsePayloads ...Payload) *application {
	if !r.enabled {
		return r
	}

	stack := NewStacktrace()

	var payloadsMap []Payload

	for _, payload := range ResponsePayloads {
		payload.Origin = map[string]string {
			"file": stack.Frames[0].AbsPath,
			"line_number": strconv.Itoa(stack.Frames[0].Lineno),
		}

		payloadsMap = append(payloadsMap, payload)
	}

	applicationConfig.sentPayloads = payloadsMap

	requestPayload := request{
		Uuid: r.Uuid(),
		Payloads: payloadsMap,
		Meta: map[string]string{
			"ray_package_version": "0.0.3",
		},
	}

	client := NewClient("http://"+r.Host()+":"+strconv.Itoa(r.Port()))

	_, err := client.sent(requestPayload)

	//Handle Error
	if err != nil {
		panic("Couldn't connect to Ray It doesn't seem to be running at " + Ray().Host() + ":" + strconv.Itoa(Ray().Port()))
	}

	return r
}
