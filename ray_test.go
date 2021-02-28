package ray

import (
	"github.com/bradleyjkemp/cupaloy/v2"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Ray().Disable()
}

func Test_CanSendAStringToRay(t *testing.T) {
	ray := Ray("hey")

	result, _ := ray.SentJsonPayloads()

	err := cupaloy.Snapshot(result)

	if err != nil {
		t.Fatal("Tests in different packages are independent of each other", err)
	}
}

func Test_CanSendAnArrayToRay(t *testing.T) {
	ray := Ray([]int{2,4,6}).Color("green")

	result, _ := ray.SentJsonPayloads()

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

//func testSnapshotsRayPayloads(ray *application, t *testing.T)  {
//	result, _ := ray.SentJsonPayloads()
//
//	err := cupaloy.Snapshot(result)
//
//	if err != nil {
//		t.Fatal("Tests in different packages are independent of each other", err)
//	}
//}