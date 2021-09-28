package main

import (
	"go-study/msgpack"
	"go-study/tools"
	"time"
)

type tTime struct {
	time.Time
}
type testStruct struct {
	A int
	T time.Time
}

type rootStruct struct {
	S testStruct
}

func main() {
	tools.PrintTypesInPackage("time")
	//tools.PrintStructsInPackage("time")
	msgpack.V5Test()
}
