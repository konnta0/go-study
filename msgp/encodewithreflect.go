package msgp

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

	fmt.Printf("%#v\n", rs)
	ebuf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(ebuf)
	if err := enc.Encode(&rs); err != nil {
		panic(err)
	}

	r := bytes.NewReader(ebuf.Bytes())
	dec := msgpack.NewDecoder(r)
	msgpack.UnregisterExt(-1)

	sl := []interface{}{
		RefStruct{},
	}
	typ := reflect.TypeOf(sl[0])
	rv := reflect.New(typ)
	itf := rv.Interface()

	if err := dec.Decode(&itf); err != nil {
		panic(err)
	}

	req := &RefStruct{Name: "bob", Age: 30}
	fmt.Printf("%#v\n", req)
	fmt.Printf("%#v\n", itf)

	j, _ := json.Marshal(itf)
	fmt.Printf("%v \n", string(j))
}
