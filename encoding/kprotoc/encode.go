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
func (e *Encoder) Encode(v interface{}, filed int) ([]byte, error) {
	obj := reflect.ValueOf(v)
	data := []byte{}
	// todo
	b, err := EncodeVarint(GetTag(filed, GetType(obj)))
	if err != nil {
		return nil, err
	}
	data = append(data, b...)
	switch GetType(obj.Interface()) {
	case Varint:
		// todo varint
		d, err := EncodeVarint(obj.Interface())
		if err != nil {
			return nil, err
		}
		data = append(data, d...)
		return data, nil
	case Bit64:
		//todo bit64
		b, err := EncodeBit64(obj.Interface())
		if err != nil {
			return nil, err
		}
		data = append(data, b...)
		return data, nil
	case StartGroup:
		//todo startGroup
		goto DecodeError
	case EndGroup:
		//todo endgroup
		goto DecodeError
	case Bit32:
		//todo bit32
		b, err := EncodeBit32(obj.Interface())
		if err != nil {
			return nil, err
		}
		data = append(data, b...)
		return data, nil
	case Struct:
		// todo struct
		le := e.getLength(obj.Interface())
		ll, err := EncodeVarint(le)
		if err != nil {
			return nil, err
		}
		data = append(data, ll...)
		for i := 0; i < obj.Elem().NumField(); i++ {
			filed++
			b, err := e.Encode(obj.Elem().Field(i).Interface(), filed)
			if err != nil {
				return nil, err
			}
			data = append(data, b...)
		}
		return data, nil
	case LengthDelimited:
		//todo encode string
		b, err := EncodeLengthDelimited(obj)
		if err != nil {
			return nil, err
		}
		ll, err := EncodeVarint(len(b))
		if err != nil {
			return nil, err
		}
		data = append(data, ll...)
		data = append(data, b...)
		return data, nil
	}
	return data, nil
DecodeError:
	return nil, errors.New("not support operation to change")

}

// getLength: dfs to get length
func (e *Encoder) getLength(v interface{}) int {
	obj := reflect.ValueOf(v)
	switch GetType(obj.Interface()) {
	case Varint:
		// todo varint
		b, _ := EncodeVarint(obj.Interface())
		return len(b)
	case Bit64:
		//todo bit64
		b, _ := EncodeBit64(obj.Interface())
		return len(b)
	case StartGroup:
		//todo startGroup
		return 0
	case EndGroup:
		//todo endgroup
		return 0
	case Bit32:
		//todo bit32
		b, _ := EncodeBit32(obj.Interface())
		return len(b)
	case Struct:
		// todo struct
		sum := 0
		for i := 0; i < obj.Elem().NumField(); i++ {
			sum += e.getLength(obj.Elem().Field(i).Interface())
		}
		return sum
	case LengthDelimited:
		// 1 byte = 8 bits
		return len(v.(string))
	default:
		return 0
	}
}
