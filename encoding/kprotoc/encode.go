package kprotoc

import "io"

type Encode struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encode {
	return &Encode{w: w}
}
