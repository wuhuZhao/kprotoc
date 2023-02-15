package kprotoc

import (
	"errors"
	"reflect"
)

type Type int

const (
	Varint = iota
	Bit64
	LengthDelimited
	StartGroup
	EndGroup
	Bit32
	Struct
)

type Field struct {
	fieldNumber int
	writeType   Type
	length      int
	value       interface{}
}

// EncodeVarint: encode varint type to []byte
func EncodeVarint(v interface{}) ([]byte, error) {
	p := reflect.TypeOf(v)
	data := []byte{}
	var fill func(real int)
	fill = func(real int) {
		for real > 0 {
			if real >= 1<<7 {
				data = append(data, byte((1<<7)|(real&((1<<7)-1))))
			} else {
				data = append(data, byte(real&((1<<7)-1)))
			}
			real >>= 7
		}
	}
	var fillN func(real uint)
	fillN = func(real uint) {
		for real > 0 {
			if real >= 1<<7 {

				data = append(data, byte((1<<7)|(real&((1<<7)-1))))
			} else {
				data = append(data, byte(real&((1<<7)-1)))
			}
			real >>= 7
		}
	}
	switch p.Kind() {
	case reflect.Int32:
		real, ok := v.(int32)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Int32")
		}
		fill(int(real))
	case reflect.Int64:
		// todo encode
		real, ok := v.(int64)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Int64")
		}
		fill(int(real))
	case reflect.Bool:
		//todo encode
		real, ok := v.(bool)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Bool")
		}
		if real {
			fill(1)
		} else {
			fill(0)
		}
	case reflect.Int16:
		//todo encode
		real, ok := v.(int16)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Int16")
		}
		fill(int(real))
	case reflect.Int:
		// todo encode
		real, ok := v.(int)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Int")
		}
		fill(real)
	case reflect.Int8:
		// todo encode
		real, ok := v.(int8)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Int8")
		}
		fill(int(real))
	case reflect.Uint:
		// todo encode
		real, ok := v.(uint)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Uint")
		}
		fillN(real)
	case reflect.Uint8:
		// todo encode
		real, ok := v.(uint8)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Uint8")
		}
		fillN(uint(real))
	case reflect.Uint16:
		// todo encode
		real, ok := v.(uint16)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Uint16")
		}
		fillN(uint(real))
	case reflect.Uint32:
		// todo encode
		real, ok := v.(uint32)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Uint32")
		}
		fillN(uint(real))
	case reflect.Uint64:
		// todo encode
		real, ok := v.(uint64)
		if !ok {
			return nil, errors.New("p.kind() can't not cast to Uint64")
		}
		fillN(uint(real))
	}
	return data, nil
}

// GetType: get v's type to encode
func GetType(v interface{}) Type {
	p := reflect.ValueOf(v).Addr().Elem()
	switch p.Kind() {
	case reflect.Int32, reflect.Int64, reflect.Uint32, reflect.Uint64, reflect.Bool, reflect.Int8, reflect.Int16, reflect.Uint16, reflect.Uint8:
		return Varint
	case reflect.Float64:
		return Bit64
	case reflect.String:
		return LengthDelimited
	case reflect.Float32:
		return Bit32
	}
	return Struct
}

// GetTag: get tag
func GetTag(fieldNum int, t Type) int {
	return fieldNum<<3 | getWriteType(t)
}

// getWriteType: get t's type encode
func getWriteType(t Type) int {
	if t == Varint {
		return 0
	} else if t == Bit64 {
		return 1
	} else if t == Bit32 {
		return 5
	} else {
		return 2
	}
}

// NewFiled: create a filed
func NewFiled(fieldNumber int, writeType Type, length int, value interface{}) *Field {
	return &Field{fieldNumber: fieldNumber, writeType: writeType, length: length, value: value}
}

// GetFieldNumber: get filed number
func (f *Field) GetFieldNumber() int {
	return f.fieldNumber
}

// GetWriterType: get writer type
func (f *Field) GetWriterType() Type {
	return f.writeType
}

// GetLength: get length
func (f *Field) GetLength() int {
	return f.length
}

// GetValue: get value
func (f *Field) GetValue() interface{} {
	return f.value
}
