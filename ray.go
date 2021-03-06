package ray

import (
	"encoding/json"
	"github.com/google/uuid"
	ray "github.com/octoper/go-ray/payloads"
	"github.com/octoper/go-ray/utils"
	"os"
	"strconv"
	"time"
)

type Callable = func() bool

type application struct {
	uuid         string
	host         string
	port         int
	enabled      bool
	sentPayloads []ray.Payload
}

type request struct {
	Uuid     string            `json:"uuid"`
	Payloads interface{}       `json:"payloads"`
	Meta     map[string]string `json:"meta"`
}

var applicationConfig = application{
	host:    "127.0.0.1",
	port:    23517,
	enabled: true,
}

// Create New Ray instance
func NewRay() *application {
	app := applicationConfig
	return &app
}

// Get the UUID
func (r *application) Uuid() string {
	return r.uuid
}

// Get the port application is running
func (r *application) Port() int {
	return r.port
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
	var payloadsMap []ray.Payload

	for _, payload := range values {
		switch payload.(type) { // nolint:gosimple
		case bool:
			payloadsMap = append(payloadsMap, ray.NewBoolPayload(payload.(bool)))
		case nil:
			payloadsMap = append(payloadsMap, ray.NewNullPayload())
		case string, int, int32, int64, float32, float64, complex64, complex128, uint, uint8, uint16, uint32, uint64:
			payloadsMap = append(payloadsMap, ray.NewCustomPayload(payload, ""))
		default:
			payloadsMap = append(payloadsMap, ray.NewDumpPayload(payload))
		}
	}

	return r.SendRequest(payloadsMap...)
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

// Get Sent Payloads as Json
func (r *application) SentJsonPayloads() ([]byte, error) {
	return json.Marshal(applicationConfig.sentPayloads)
}

// Enable sending payloads to Ray
func (r *application) Enable() {
	applicationConfig.enabled = true
}

// Prevents sending payloads to Ray
func (r *application) Disable() {
	applicationConfig.enabled = false
}

// Check if Ray is enabled
func (r *application) Enabled() bool {
	return applicationConfig.enabled
}

// Check if Ray is disabled
func (r *application) Disabled() bool {
	return !applicationConfig.enabled
}

// Create New Screen
func (r *application) NewScreen(name string) *application {
	return r.SendRequest(ray.NewNewScreenPayload(name))
}

// Color
func (r *application) Color(color string) *application {
	return r.SendRequest(ray.NewColorPayload(color))
}

// Send custom payload
func (r *application) SendCustom(content interface{}, label string) *application {
	return r.SendRequest(ray.NewCustomPayload(content, label))
}

// Size
func (r *application) Size(size string) *application {
	return r.SendRequest(ray.NewSizePayload(size))
}

// Hide
func (r *application) Hide() *application {
	return r.SendRequest(ray.NewHidePayload())
}

// Hide App
func (r *application) HideApp() *application {
	return r.SendRequest(ray.NewHideAppPayload())
}

// Show App
func (r *application) ShowApp() *application {
	return r.SendRequest(ray.NewShowAppPayload())
}

// Clear Screen
func (r *application) ClearScreen() *application {
	return r.SendRequest(ray.NewClearScreenPayload())
}

// Clear All
func (r *application) ClearAll() *application {
	return r.SendRequest(ray.NewClearAllPayload())
}

// HTML
func (r *application) Html(html string) *application {
	return r.SendRequest(ray.NewHtmlPayload(html))
}

// Notify
func (r *application) Notify(text string) *application {
	return r.SendRequest(ray.NewNotifyPayload(text))
}

// Pass
func (r *application) Pass(arg interface{}) interface{} {
	r.Send(arg)
	return arg
}

// Boolean
func (r *application) Bool(bool bool) *application {
	return r.SendRequest(ray.NewBoolPayload(bool))
}

// Null
func (r *application) Null() *application {
	return r.SendRequest(ray.NewNullPayload())
}

// Charles
func (r *application) Charles() *application {
	return r.Send("🎶 🎹 🎷 🕺")
}

// String
func (r *application) String(str string) *application {
	return r.SendRequest(ray.NewStringPayload(str))
}

//Time
func (r *application) Time(time time.Time) *application {
	return r.SendRequest(ray.NewTimePayload(time, "2006-01-02 15:04:05"))
}

// Time with Format
func (r *application) TimeWithFormat(time time.Time, format string) *application {
	return r.SendRequest(ray.NewTimePayload(time, format))
}

/**
 * Sends the provided value(s) encoded as a JSON string using json.Marshal.
 */
func (r *application) ToJson(jsons ...interface{}) *application {
	for _, jsonValue := range jsons {
		r.SendRequest(ray.NewJsonStringPayload(jsonValue))
	}
	return r
}

// Image
func (r *application) Image(value string) *application {
	return r.SendRequest(ray.NewImagePayload(value))
}

// Ban
func (r *application) Ban() *application {
	return r.Send("🕶")
}

// Die
func (r *application) Die() {
	r.DieWithStatusCode(1)
}

// Die with Status Code
func (r *application) DieWithStatusCode(status int) {
	os.Exit(status)
}

// Remove
func (r *application) Remove() *application {
	return r.SendRequest(ray.NewRemovePayload())
}

// Show When
func (r *application) ShowWhen(show interface{}) *application {
	switch show.(type) { // nolint:gosimple
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
	switch show.(type) { // nolint:gosimple
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
func (r *application) SendRequest(ResponsePayloads ...ray.Payload) *application {
	var payloadsMap []ray.Payload

	stack := utils.NewStacktrace()
	stackAbsPath := ""
	stackLineNo := ""

	if len(stack.Frames) > 0 {
		stackAbsPath = stack.Frames[0].AbsPath
		stackLineNo = strconv.Itoa(stack.Frames[0].Lineno)
	}

	for _, payload := range ResponsePayloads {
		payload.Origin = map[string]string{
			"file":        stackAbsPath,
			"line_number": stackLineNo,
		}

		payloadsMap = append(payloadsMap, payload)
	}

	applicationConfig.sentPayloads = payloadsMap

	requestPayload := request{
		Uuid:     r.Uuid(),
		Payloads: payloadsMap,
		Meta: map[string]string{
			"ray_package_version": "0.0.3",
		},
	}

	if !r.enabled {
		return r
	}

	client := utils.NewClient("http://" + r.Host() + ":" + strconv.Itoa(r.Port()))

	_, err := client.Sent(requestPayload)

	//Handle Error
	if err != nil {
		panic("Couldn't connect to Ray It doesn't seem to be running at " + Ray().Host() + ":" + strconv.Itoa(Ray().Port()))
	}

	return r
}
