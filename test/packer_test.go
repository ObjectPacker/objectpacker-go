package main

import (
	"../pack"
	"testing"
)
//
type X struct {
	A float64
	B float64
}
//
func TestPack(t *testing.T) {

	p := X{B: 20}

	b := []byte(`
{"A":99,"B":102}
	`)
	pack.Packer(&b, &p)

	if p.B != 102{
		t.Error("test fail")
	}

}
