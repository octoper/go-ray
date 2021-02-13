package utils

import (
	"runtime"
)

type Origin struct {
	file string
	line int
}

func GetBackTrace(depth int) (string, int) {
	_, fn, line, _ := runtime.Caller(depth)

	return fn, line
}