// Generated by: main
// TypeWriter: holder
// Directive: +gen on FooInner

package ex

import (
	"github.com/tendermint/go-wire/data"
)

// Auto-generated adapters for happily unmarshaling interfaces
// Apache License 2.0
// Copyright (c) 2017 Ethan Frey (ethan.frey@tendermint.com)

type Foo struct {
	FooInner
}

var FooMapper = data.NewMapper(Foo{})

func (h Foo) MarshalJSON() ([]byte, error) {
	return FooMapper.ToJSON(h.FooInner)
}

func (h *Foo) UnmarshalJSON(data []byte) (err error) {
	parsed, err := FooMapper.FromJSON(data)
	if err == nil && parsed != nil {
		h.FooInner = parsed.(FooInner)
	}
	return err
}

// Unwrap recovers the concrete interface safely (regardless of levels of embeds)
func (h Foo) Unwrap() FooInner {
	hi := h.FooInner
	for wrap, ok := hi.(Foo); ok; wrap, ok = hi.(Foo) {
		hi = wrap.FooInner
	}
	return hi
}

func (h Foo) Empty() bool {
	return h.FooInner == nil
}

/*** below are bindings for each implementation ***/

func init() {
	FooMapper.RegisterImplementation(Bling{}, "bb", 0x1)
}

func (hi Bling) Wrap() Foo {
	return Foo{hi}
}

func init() {
	FooMapper.RegisterImplementation(&Fuzz{}, "ff", 0x2)
}

func (hi *Fuzz) Wrap() Foo {
	return Foo{hi}
}
