package msgp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/ugorji/go/codec"
	"github.com/vmihailenco/msgpack/v5"
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

func V5Test() {
	fmt.Println("hello world v5")
	msgpack.Register(time.Time{}, timeEncodeFunc, timeDecoderFunc)

	s := rootStruct{
		S: testStruct{
			A: 1234,
			T: time.Now(),
		},
	}
	ebuf := &bytes.Buffer{}
	enc := msgpack.NewEncoder(ebuf)
	if err := enc.Encode(&s); err != nil {
		panic(err)
	}

	var item rootStruct
	r := bytes.NewReader(ebuf.Bytes())
	dec := msgpack.NewDecoder(r)
	msgpack.UnregisterExt(-1)
	if err := dec.Decode(&item); err != nil {
		panic(err)
	}

	j, _ := json.Marshal(item)
	fmt.Printf("%v \n", string(j))
}

func timeEncodeFunc(e *msgpack.Encoder, v reflect.Value) error {
	fmt.Printf("encode time\n")
	// from https://github.com/vmihailenco/msgpack/blob/d72feb0678b758964b52d8aa092e3eb2962f79c7/time.go#L20
	return e.EncodeTime(v.Interface().(time.Time))
}

func timeDecoderFunc(d *msgpack.Decoder, v reflect.Value) error {
	fmt.Printf("decode time\n")
	// from https://github.com/vmihailenco/msgpack/blob/d72feb0678b758964b52d8aa092e3eb2962f79c7/time.go#L23
	tm, err := d.DecodeTime()
	if err != nil {
		return err
	}
	v.Set(reflect.ValueOf(tm))
	return nil
}

func codecTest() {
	fmt.Println("hello world codec")

	var mh codec.MsgpackHandle

	s := rootStruct{
		S: testStruct{
			A: 1234,
			T: time.Now(),
		},
	}
	var h = &mh

	ebuf := &bytes.Buffer{}
	enc := codec.NewEncoder(ebuf, h)
	if err := enc.Encode(s); err != nil {
		fmt.Printf("encode err %v\n", err)
	}

	var actual rootStruct
	r := bytes.NewReader(ebuf.Bytes())
	dec := codec.NewDecoder(r, h)
	if err := dec.Decode(&actual); err != nil {
		fmt.Printf("decode err %v\n", err)
	}

	j, _ := json.Marshal(actual)
	fmt.Printf("%v \n", string(j))
}
