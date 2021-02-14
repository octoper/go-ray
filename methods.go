package ray

import (
	"github.com/octoper/ray/payloads"
	"os"
	"time"
)

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

// Json String
func (r *Application) Json(json string) *Application {
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