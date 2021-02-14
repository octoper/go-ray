package ray

import "testing"

func Test_Ray(t *testing.T) {
	Ray("this is green").Green()
	Ray("this is orange").Orange()
	Ray("this is red").Red()
}