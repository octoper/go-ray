package utils

import (
	"runtime"
)

type Origin struct {
	file string
	line int
}

func GetBackTrace() (string, int) {
	_, fn, line, _ := runtime.Caller(3)

	return fn, line
}