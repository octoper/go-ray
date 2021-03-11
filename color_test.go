package ray

import (
	"github.com/bradleyjkemp/cupaloy/v2"
	"testing"
)

func TestApplication_Red(t *testing.T) {
	// t.Parallel()

	ray := Ray().Red()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}


func TestApplication_Green(t *testing.T) {
	// t.Parallel()

	ray := Ray().Green()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Blue(t *testing.T) {
	// t.Parallel()

	ray := Ray().Blue()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Purple(t *testing.T) {
	// t.Parallel()

	ray := Ray().Purple()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Orange(t *testing.T) {
	// t.Parallel()

	ray := Ray().Orange()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

func TestApplication_Gray(t *testing.T) {
	// t.Parallel()

	ray := Ray().Gray()

	result, _ := ray.SentJsonPayloads()

	cupaloy.SnapshotT(t, result)
}

