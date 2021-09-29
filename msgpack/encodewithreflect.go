package msgpack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/vmihailenco/msgpack/v5"
)

type RefStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func EncodeWithReflect() {

	rs := RefStruct{
		Name: "jon",
		Age:  20,
	}

	ebuf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(ebuf)
	if err := enc.Encode(&rs); err != nil {
		panic(err)
	}

	typ := reflect.TypeOf(new(RefStruct))
	rv := reflect.New(typ)
	itf := rv.Interface()

	r := bytes.NewReader(ebuf.Bytes())
	dec := msgpack.NewDecoder(r)
	msgpack.UnregisterExt(-1)
	if err := dec.Decode(&itf); err != nil {
		panic(err)
	}

	j, _ := json.Marshal(itf)
	fmt.Printf("%v \n", string(j))
}
