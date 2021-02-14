package ray

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/octoper/go-ray/payloads"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Callable func() bool

type Application struct {
	uuid string
	host string
	port int
}

type Request struct {
	Uuid     string      `json:"uuid"`
	Payloads interface{} `json:"payloads"`
	Meta     map[string]string     `json:"meta"`
}

var ApplicationConfig = Application {
	host: "127.0.0.1",
	port: 23517,
}

// Create New Ray instance
func NewRay() *Application {
	ray := ApplicationConfig
	return &ray
}

// Get the UUID
func (r *Application) Uuid() string {
	return r.uuid
}

// Get the port Application is running
func (r *Application) Port() int {
	return r.port
}

// Set the port Application is running
func (r *Application) SetPort(port int) {
	ApplicationConfig.port = port
}

// Get the host Application is running
func (r *Application) Host() string {
	return r.host
}

// Set the host Application is running
func (r *Application) SetHost(host string) {
	ApplicationConfig.host = host
}


func Ray(values ...interface{}) *Application {
	r := NewRay()

	r.uuid = uuid.New().String()

	if values != nil {
		r.Send(values...)
	}

	return r
}

// Send Values
func (r *Application) Send(values ...interface{}) *Application {
	var payloadsMap []payloads.Payload

	for _, payload := range values {
		switch payload.(type) {
		case bool:
			payloadsMap = append(payloadsMap, payloads.NewBoolPayload(payload.(bool)))
		case nil:
			payloadsMap = append(payloadsMap, payloads.NewNullPayload())
		default:
			payloadsMap = append(payloadsMap, payloads.NewLogPayload(payload))
		}
	}

	return r.SendRequest(payloadsMap...)
}

// Create New Screen
func (r *Application) NewScreen(name string) *Application {
	return r.SendRequest(payloads.NewNewScreenPayload(name))
}

// Color
func (r *Application) Color(color string) *Application {
	return r.SendRequest(payloads.NewColorPayload(color))
}

// Send custom payload
func (r *Application) SendCustom(content interface{}, label string) *Application {
	return r.SendRequest(payloads.NewCustomPayload(content, label))
}

// Size
func (r *Application) Size(size string) *Application {
	return r.SendRequest(payloads.NewSizePayload(size))
}

// Remove
func (r *Application) Remove() *Application {
	return r.SendRequest(payloads.NewRemovePayload())
}

// Hide
func (r *Application) Hide() *Application {
	return r.SendRequest(payloads.NewHidePayload())
}

// Hide App
func (r *Application) HideApp() *Application {
	return r.SendRequest(payloads.NewHideAppPayload())
}

// Show App
func (r *Application) ShowApp() *Application {
	return r.SendRequest(payloads.NewShowAppPayload())
}

// Clear Screen
func (r *Application) ClearScreen() *Application {
	return r.SendRequest(payloads.NewClearScreenPayload())
}

// Clear All
func (r *Application) ClearAll() *Application {
	return r.SendRequest(payloads.NewClearAllPayload())
}

// HTML
func (r *Application) Html(html string) *Application {
	return r.SendRequest(payloads.NewHtmlPayload(html))
}

// Notify
func (r *Application) Notify(text string) *Application {
	return r.SendRequest(payloads.NewNotifyPayload(text))
}

// Pass
func (r *Application) Pass(arg interface{}) interface{} {
	r.Send(arg)
	return arg
}

// Boolean
func (r *Application) Bool(bool bool) *Application {
	return r.SendRequest(payloads.NewBoolPayload(bool))
}

// Null
func (r *Application) Null() *Application {
	return r.SendRequest(payloads.NewNullPayload())
}

// Charles
func (r *Application) Charles() *Application {
	return r.Send("🎶 🎹 🎷 🕺")
}

// String
func (r *Application) String(str string) *Application {
	return r.SendRequest(payloads.NewStringPayload(str))
}

// Time
func (r *Application) Time(time time.Time) *Application {
	return r.SendRequest(payloads.NewTimePayload(time, "2021-02-13 18:38:20"))
}

// Json
func (r *Application) Json(json interface{}) *Application {
	return r.SendRequest(payloads.NewJsonStringPayload(json))
}

// Image
func (r *Application) Image(value string) *Application {
	return r.SendRequest(payloads.NewImagePayload(value))
}

// Ban
func (r *Application) Ban() *Application {
	return r.Send("🕶")
}

// Die
func (r *Application) Die() {
	os.Exit(1)
}

// Show When
func (r *Application) ShowWhen(show interface{}) *Application {
	switch show.(type) {
	case Callable:
		show = show.(Callable)
	}

	if !show.(bool) {
		return r.Remove()
	}

	return r
}

// Show If
func (r *Application) ShowIf(show interface{}) *Application {
	return r.ShowWhen(show)
}

// Remove When
func (r *Application) RemoveWhen(show interface{}) *Application {
	switch show.(type) {
	case Callable:
		show = show.(Callable)
	}

	if show.(bool) {
		return r.Remove()
	}

	return r
}

// Remove If
func (r *Application) RemoveIf(show interface{}) *Application {
	return r.RemoveWhen(show)
}

// Set the host Application is running
func (r *Application) SendRequest(ResponsePayloads ...payloads.Payload) *Application {
	//file, line := utils.GetBackTrace(4)

	var payloadsMap []payloads.Payload

	for _, payload := range ResponsePayloads {
		payload.Origin = map[string]string {
			"file": "",
			"line_number": "",
		}

		payloadsMap = append(payloadsMap, payload)
	}

	requestPayload := Request{
		Uuid: r.Uuid(),
		Payloads: payloadsMap,
		Meta: map[string]string{
			"ray_package_version": "0.0.1",
		},
	}

	//
	//
	//fmt.Println(file)
	//fmt.Println(line)

	requestJson, _ := json.Marshal(requestPayload)

	//fmt.Println(string(requestJson))

	responseBody := bytes.NewBuffer(requestJson)

	//Make a request to Ray
	resp, err := http.Post("http://"+ r.Host() +":"+ strconv.Itoa(r.Port()), "application/json", responseBody)

	//Handle Error
	if err != nil {
		panic("Couldn't connect to Ray It doesn't seem to be running at " + r.Host() + ":" + strconv.Itoa(r.Port()))
	}

	defer resp.Body.Close()

	return r
}
