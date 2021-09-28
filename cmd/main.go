package main

import (
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
	tools.SearchStructsInPackage("time")
	//tools.PrintTypesInPackage("time")
	//tools.PrintStructsInPackage("time")
	//msgpack.V5Test()
}
