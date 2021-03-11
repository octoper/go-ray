package main

import (
	"time"
	"github.com/octoper/go-ray"
)

func main() {
	ray.Ray().ClearAll()
	myRay := ray.Ray(4)
	for i := 0; i < 10; i++ {
		time.Sleep(1)
		myRay.Send(2)
	}	
}