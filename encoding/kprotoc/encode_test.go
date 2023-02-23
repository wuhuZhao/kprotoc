package kprotoc

import (
	"reflect"
	"testing"
)

type TestData struct {
	a int
	b string
}

func TestEncode(t *testing.T) {
	e := NewEncoder(nil)
	d := &TestData{a: 100, b: "test"}
	t.Logf("%v\n", reflect.ValueOf(d).CanAddr())
	data, err := e.Encode(d, 1)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	t.Logf("%v\n", data)
}
