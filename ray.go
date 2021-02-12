package ray

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"ray/payloads"
)

type Callable func() bool

var rayHost = "127.0.0.1"
var rayPort = 23517

type Ray struct {
	uuid string
	host string
	port int
}

type Request struct {
	Uuid     string      `json:"uuid"`
	Payloads interface{} `json:"payloads"`
	Meta     string      `json:"meta"`
}

func ray(values ...interface{}) *Ray {
	r := NewRay()

	if values != nil {
		r.Send(values...)
	}

	return r
}

func NewRay() *Ray {
	ray := Ray{
		uuid: uuid.New().String(),
		host: rayHost,
		port: rayPort,
	}
	return &ray
}

// Send Values
func (r *Ray) Send(values ...interface{}) *Ray {
	var payloadsMap []interface{}

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

	return r.SendRequest(payloadsMap)
}

// Get the UUID
func (r *Ray) Uuid() string {
	return r.uuid
}

// Get the port Ray is running
func (r *Ray) Port() int {
	return r.port
}

// Set the port Ray is running
func (r *Ray) SetPort(port int) {
	rayPort = port
}

// Get the host Ray is running
func (r *Ray) Host() string {
	return r.host
}

// Set the host Ray is running
func (r *Ray) SetHost(host string) {
	rayHost = host
}

// Create New Screen
func (r *Ray) NewScreen(name string) *Ray {
	return r.SendRequest(payloads.NewNewScreenPayload(name))
}

// Color
func (r *Ray) Color(color string) *Ray {
	return r.SendRequest(payloads.NewColorPayload(color))
}

// Send custom payload
func (r *Ray) SendCustom(content interface{}, label string) *Ray {
	return r.SendRequest(payloads.NewCustomPayload(content, label))
}

// Size
func (r *Ray) Size(size string) *Ray {
	return r.SendRequest(payloads.NewSizePayload(size))
}

// Remove
func (r *Ray) Remove() *Ray {
	return r.SendRequest(payloads.NewRemovePayload())
}

// Hide
func (r *Ray) Hide() *Ray {
	return r.SendRequest(payloads.NewHidePayload())
}

// Hide App
func (r *Ray) HideApp() *Ray {
	return r.SendRequest(payloads.NewHideAppPayload())
}

// Show App
func (r *Ray) ShowApp() *Ray {
	return r.SendRequest(payloads.NewShowAppPayload())
}

// Clear Screen
func (r *Ray) ClearScreen() *Ray {
	return r.SendRequest(payloads.NewClearScreenPayload())
}

// Clear All
func (r *Ray) ClearAll() *Ray {
	return r.SendRequest(payloads.NewClearAllPayload())
}

// HTML
func (r *Ray) Html(html string) *Ray {
	return r.SendRequest(payloads.NewHtmlPayload(html))
}

// Notify
func (r *Ray) Notify(text string) *Ray {
	return r.SendRequest(payloads.NewNotifyPayload(text))
}

// Pass
func (r *Ray) Pass(arg interface{}) *Ray {
	return r.Send(arg)
}

// Boolean
func (r *Ray) Bool(bool bool) *Ray {
	return r.SendRequest(payloads.NewBoolPayload(bool))
}

// Null
func (r *Ray) Null() *Ray {
	return r.SendRequest(payloads.NewNullPayload())
}

// Charles
func (r *Ray) Charles() *Ray {
	return r.Send("🎶 🎹 🎷 🕺")
}

// String
func (r *Ray) String(str string) *Ray {
	return r.SendRequest(payloads.NewStringPayload(str))
}

// Json String
func (r *Ray) Json(json string) *Ray {
	return r.SendRequest(payloads.NewJsonStringPayload(json))
}

// Ban
func (r *Ray) Ban() *Ray {
	return r.Send("🕶")
}

// Die
func (r *Ray) Die() {
	os.Exit(1)
}

// Show When
func (r *Ray) ShowWhen(show interface{}) *Ray {
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
func (r *Ray) ShowIf(show interface{}) *Ray {
	return r.ShowWhen(show)
}

// Remove When
func (r *Ray) RemoveWhen(show interface{}) *Ray {
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
func (r *Ray) RemoveIf(show interface{}) *Ray {
	return r.RemoveWhen(show)
}

// Set the host Ray is running
func (r *Ray) SendRequest(payloads interface{}) *Ray {
	requestPayload := Request{
		Uuid:     r.Uuid(),
		Payloads: payloads,
		Meta:     "",
	}

	requestJson, _ := json.Marshal(requestPayload)

	fmt.Println(string(requestJson))

	//req, _ := http.NewRequest("POST", "http://" + r.Host() +":"+ string(r.Port()), bytes.NewBuffer(requestJson))
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	return r
}
