package ray

import (
	"github.com/bradleyjkemp/cupaloy/v2"
	"testing"
)

func Test_CanSendAStringToRay(t *testing.T) {
	ray := Ray("hey")

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func Test_CanSendAnArrayToRay(t *testing.T) {
	ray := Ray([]int{2,4,6}).Color("green")

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Bool(t *testing.T) {

	ray := Ray().Bool(false)

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Ban(t *testing.T) {

	ray := Ray().Ban()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Color(t *testing.T) {
	t.Parallel()

	ray := Ray().Color("green")

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Charles(t *testing.T) {
	t.Parallel()

	ray := Ray().Charles()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}