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
	switch GetType(obj) {
	case Varint:
		// todo varint
		d, err := EncodeVarint(obj)
		if err != nil {
			return nil, err
		}
		data = append(data, d...)
		return data, nil
	case Bit64:
		//todo bit64
		b, err := EncodeBit64(obj)
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
		b, err := EncodeBit32(obj)
		if err != nil {
			return nil, err
		}
		data = append(data, b...)
		return data, nil
	case Struct:
		// todo struct
		b, err := e.Encode(obj, filed+1)
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
	case LengthDelimited:
		//todo encode string
		return data, nil
	}
	return data, nil
DecodeError:
	return nil, errors.New("not support operation to change")

}
