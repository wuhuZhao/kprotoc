package kprotoc

import (
	"io"
	"reflect"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode: bfs to encode the data
func (e *Encoder) Encode(v interface{}) error {
	obj := reflect.ValueOf(v).Addr().Elem()
	data := []byte{}
	eles := []reflect.Value{obj}
	for len(eles) > 0 {
		for i := 0; i < len(eles); i++ {
			obj := eles[i]
			switch GetType(obj) {
			case Varint:
				// todo varint
				data = append(data, EncodeVarint(GetTag(i+1, GetType(obj)))...)
			case Bit64:
				//todo bit64
			case StartGroup:
				//todo startGroup
			case EndGroup:
				//todo endgroup
			case Bit32:
				//todo bit32
			case Struct:
				// todo struct
			}
		}
	}
}

// getLength: dfs to get length
func (e *Encoder) getLength(v interface{}) int {
	obj := reflect.ValueOf(v).Addr().Elem()
	switch GetType(obj) {
	case Varint:
		// todo varint
		return 0
	case Bit64:
		//todo bit64
		return 0
	case StartGroup:
		//todo startGroup
		return 0
	case EndGroup:
		//todo endgroup
		return 0
	case Bit32:
		//todo bit32
		return 0
	case Struct:
		// todo struct
		sum := 0
		for i := 0; i < obj.NumField(); i++ {
			sum += e.getLength(obj.Field(i))
		}
		return sum
	case LengthDelimited:
		// 1 byte = 8 bits
		return len(v.(string)) * 8
	}
	return e.getLength(v)
}
