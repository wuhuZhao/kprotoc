package kprotoc

import "io"

type Decode struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decode {
	return &Decode{r: r}
}
