package kprotoc

import (
	"errors"
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
func (e *Encoder) Encode(v interface{}) ([]byte, error) {
	obj := reflect.ValueOf(v).Addr().Elem()
	data := []byte{}
	eles := []reflect.Value{obj}
	for len(eles) > 0 {
		for i := 0; i < len(eles); i++ {
			obj := eles[i]
			b, err := EncodeVarint(GetTag(i+1, GetType(obj)))
			if err != nil {
				return nil, err
			}
			data = append(data, b...)
			switch GetType(obj) {
			case Varint:
				// todo varint
				d, err := EncodeVarint(obj)
				if err != nil {
					return nil, err
				}
				data = append(data, d...)
				continue
			case Bit64:
				//todo bit64
				b, err := EncodeBit64(obj)
				if err != nil {
					return nil, err
				}
				data = append(data, b...)
				continue
			case StartGroup:
				//todo startGroup
				goto DecodeError
			case EndGroup:
				//todo endgroup
				goto DecodeError
			case Bit32:
				//todo bit32
				b, err := EncodeBit32(obj)
				if err != nil {
					return nil, err
				}
				data = append(data, b...)
				continue
			case Struct:
				// todo struct
				b, err := e.Encode(obj)
				if err != nil {
					return nil, err
				}
				ll, err := EncodeVarint(len(b))
				if err != nil {
					return nil, err
				}
				data = append(data, ll...)
				data = append(data, b...)
				continue
			}
		}
	}
	return data, nil
DecodeError:
	return nil, errors.New("not support operation to change")

}

// getLength: dfs to get length
// func (e *Encoder) getLength(v interface{}) int {
// 	obj := reflect.ValueOf(v).Addr().Elem()
// 	switch GetType(obj) {
// 	case Varint:
// 		// todo varint
// 		return 0
// 	case Bit64:
// 		//todo bit64
// 		return 0
// 	case StartGroup:
// 		//todo startGroup
// 		return 0
// 	case EndGroup:
// 		//todo endgroup
// 		return 0
// 	case Bit32:
// 		//todo bit32
// 		return 0
// 	case Struct:
// 		// todo struct
// 		sum := 0
// 		for i := 0; i < obj.NumField(); i++ {
// 			sum += e.getLength(obj.Field(i))
// 		}
// 		return sum
// 	case LengthDelimited:
// 		// 1 byte = 8 bits
// 		return len(v.(string)) * 8
// 	}
// 	return e.getLength(v)
// }
