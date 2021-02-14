package ray

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/octoper/ray/payloads"
)

type Callable func() bool

func Ray(values ...interface{}) *Application {
	r := NewRay()

	if values != nil {
		r.Send(values...)
	}

	return r
}

// Create New Ray instance
func NewRay() *Application {
	ray := Application{
		uuid: uuid.New().String(),
		host: rayHost,
		port: rayPort,
	}
	return &ray
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

// Get the UUID
func (r *Application) Uuid() string {
	return r.uuid
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

	fmt.Println(string(requestJson))

	//responseBody := bytes.NewBuffer(requestJson)
	////Leverage Go's HTTP Post function to make request
	//resp, err := http.Post("http://"+ r.Host() +":"+ strconv.Itoa(r.Port()), "application/json", responseBody)
	////Handle Error
	//if err != nil {
	//	log.Fatalf("An Error Occured %v", err)
	//}
	//defer resp.Body.Close()
	////Read the response body
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//sb := string(body)
	//log.Printf(sb)

	return r
}
