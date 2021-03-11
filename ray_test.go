package ray

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Ray().Disable()

	os.Exit(m.Run())
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

