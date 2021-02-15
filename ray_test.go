package ray

import (
	"github.com/bradleyjkemp/cupaloy/v2"
	"os"
	"testing"
)

func Test_CanSendAStringToRay(t *testing.T) {
	Ray().Json([]byte(`{"type":"log","content":{"values":[[2,4,6]]}`))

	result, _ := Ray("hey").SentJsonPayloads()

	if len(result) != 1 {
		t.Fatalf(`There is something wrong with the payloads`)
	}

	err := cupaloy.Snapshot(result)
	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}

func Test_CanSendAnArrayToRay(t *testing.T) {
	result, _ := Ray([]int{2,4,6}).SentJsonPayloads()

	err := cupaloy.Snapshot(result)
	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Test does not run on GitHub actions")
	}
}