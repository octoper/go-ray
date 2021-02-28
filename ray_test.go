package ray

import (
	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Ray().Disable()

	os.Exit(m.Run())
}

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


func TestApplication_SetHost(t *testing.T) {
	t.Parallel()

	Ray().SetHost("192.168.1.255")

	assert.Equal(t, "192.168.1.255", Ray().Host())
}


func TestApplication_SetPort(t *testing.T) {
	t.Parallel()

	Ray().SetPort(27154)

	assert.Equal(t, 27154, Ray().Port())
}

func TestApplication_Color(t *testing.T) {
	t.Parallel()

	ray := Ray().Color("green")

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Red(t *testing.T) {
	t.Parallel()

	ray := Ray().Red()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

