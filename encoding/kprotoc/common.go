package kprotoc

type Type int

const (
	Varint = iota
	Bit64
	LengthDelimited
	StartGroup
	EndGroup
	Bit32
)

type Field struct {
	fieldNumber int
	writeType   Type
	length      int
	value       interface{}
}

func NewFiled(fieldNumber int, writeType Type, length int, value interface{}) *Field {
	return &Field{fieldNumber: fieldNumber, writeType: writeType, length: length, value: value}
}

func (f *Field) GetFieldNumber() int {
	return f.fieldNumber
}

func (f *Field) GetWriterType() Type {
	return f.writeType
}

func (f *Field) GetLength() int {
	return f.length
}

func (f *Field) GetValue() interface{} {
	return f.value
}
