package kprotoc

import (
	"fmt"
	"testing"
)

func TestEncodeVarint(t *testing.T) {
	data, err := EncodeVarint(100)
	if err != nil {
		t.Errorf("encode error: %v\n", err)
	}
	for i := 0; i < len(data); i++ {
		fmt.Printf("%b ", uint8(data[i]))
	}
}
