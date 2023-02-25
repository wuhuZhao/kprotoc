package kprotoc

import (
	"reflect"
	"testing"
)

type TestData struct {
	A int
	B string
}

func TestEncode(t *testing.T) {
	e := NewEncoder(nil)
	d := &TestData{A: 100, B: "test"}
	data, err := e.Encode(d, 1)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	t.Logf("%v\n", data)
}

func TestFiled(t *testing.T) {
	t.Log(GetType(1))
	t.Log(reflect.ValueOf(1).Kind() == reflect.Int)
}
