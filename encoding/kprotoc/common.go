package kprotoc

import (
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
func EncodeVarint(v interface{}) []byte {
	p := reflect.TypeOf(v)
	data := []byte{}
	switch p.Kind() {
	case reflect.Int32:
		// todo encode
	case reflect.Int64:
		// todo encode
	case reflect.Bool:
		//todo encode
	case reflect.Int16:
		//todo encode
	case reflect.Int:
		// todo encode
	case reflect.Int8:
		// todo encode
	case reflect.Uint:
		// todo encode
	case reflect.Uint8:
		// todo encode
	case reflect.Uint16:
		// todo encode
	case reflect.Uint32:
		// todo encode
	case reflect.Uint64:
		// todo encode
	}
	return data
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
