package ray

import (
	"github.com/bradleyjkemp/cupaloy/v2"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Ray().Disable()

	os.Exit(m.Run())
}

func Test_CanSendAStringToRay(t *testing.T) {
	t.Parallel()

	ray := Ray("hey")

	result, _ := ray.SentJsonPayloads()

	err := cupaloy.Snapshot(result)

	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}

func Test_CanSendAnArrayToRay(t *testing.T) {
	t.Parallel()

	ray := Ray([]int{2,4,6}).Color("green")

	result, _ := ray.SentJsonPayloads()

	err := cupaloy.Snapshot(result)

	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}
