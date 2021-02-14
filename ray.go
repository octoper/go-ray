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

type callable func() bool

type application struct {
	uuid string
	host string
	port int
}

type request struct {
	Uuid     string      `json:"uuid"`
	Payloads interface{} `json:"payloads"`
	Meta     map[string]string     `json:"meta"`
}

var applicationConfig = application{
	host: "127.0.0.1",
	port: 23517,
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
func (r *application) NewScreen(name string) *application {
	return r.SendRequest(payloads.NewNewScreenPayload(name))
}

// Color
func (r *application) Color(color string) *application {
	return r.SendRequest(payloads.NewColorPayload(color))
}

// Send custom payload
func (r *application) SendCustom(content interface{}, label string) *application {
	return r.SendRequest(payloads.NewCustomPayload(content, label))
}

// Size
func (r *application) Size(size string) *application {
	return r.SendRequest(payloads.NewSizePayload(size))
}

// Remove
func (r *application) Remove() *application {
	return r.SendRequest(payloads.NewRemovePayload())
}

// Hide
func (r *application) Hide() *application {
	return r.SendRequest(payloads.NewHidePayload())
}

// Hide App
func (r *application) HideApp() *application {
	return r.SendRequest(payloads.NewHideAppPayload())
}

// Show App
func (r *application) ShowApp() *application {
	return r.SendRequest(payloads.NewShowAppPayload())
}

// Clear Screen
func (r *application) ClearScreen() *application {
	return r.SendRequest(payloads.NewClearScreenPayload())
}

// Clear All
func (r *application) ClearAll() *application {
	return r.SendRequest(payloads.NewClearAllPayload())
}

// HTML
func (r *application) Html(html string) *application {
	return r.SendRequest(payloads.NewHtmlPayload(html))
}

// Notify
func (r *application) Notify(text string) *application {
	return r.SendRequest(payloads.NewNotifyPayload(text))
}

// Pass
func (r *application) Pass(arg interface{}) interface{} {
	r.Send(arg)
	return arg
}

// Boolean
func (r *application) Bool(bool bool) *application {
	return r.SendRequest(payloads.NewBoolPayload(bool))
}

// Null
func (r *application) Null() *application {
	return r.SendRequest(payloads.NewNullPayload())
}

// Charles
func (r *application) Charles() *application {
	return r.Send("🎶 🎹 🎷 🕺")
}

// String
func (r *application) String(str string) *application {
	return r.SendRequest(payloads.NewStringPayload(str))
}

// Time
func (r *application) Time(time time.Time) *application {
	return r.SendRequest(payloads.NewTimePayload(time, "2021-02-13 18:38:20"))
}

// Json
func (r *application) Json(json interface{}) *application {
	return r.SendRequest(payloads.NewJsonStringPayload(json))
}

// Image
func (r *application) Image(value string) *application {
	return r.SendRequest(payloads.NewImagePayload(value))
}

// Ban
func (r *application) Ban() *application {
	return r.Send("🕶")
}

// Die
func (r *application) Die() {
	os.Exit(1)
}

// Show When
func (r *application) ShowWhen(show interface{}) *application {
	switch show.(type) {
	case callable:
		show = show.(callable)
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
	case callable:
		show = show.(callable)
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
func (r *application) SendRequest(ResponsePayloads ...payloads.Payload) *application {
	//file, line := utils.GetBackTrace(4)

	var payloadsMap []payloads.Payload

	for _, payload := range ResponsePayloads {
		payload.Origin = map[string]string {
			"file": "",
			"line_number": "",
		}

		payloadsMap = append(payloadsMap, payload)
	}

	requestPayload := request{
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
