package amino_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	amino "github.com/tendermint/go-amino"
)

type DCFoo1 struct{ a string }

func newDCFoo1(a string) *DCFoo1            { return &DCFoo1{a: a} }
func (dcf *DCFoo1) MarshalAmino() string    { return dcf.a }
func (dcf *DCFoo1) UnmarshalAmino(s string) { dcf.a = s }

func TestDeepCopyFoo1(t *testing.T) {
	dcf1 := newDCFoo1("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo1)
	assert.Equal(t, "foobar", dcf2.a)
}

type DCFoo2 struct{ a string }

func newDCFoo2(a string) *DCFoo2            { return &DCFoo2{a: a} }
func (dcf DCFoo2) MarshalAmino() string     { return dcf.a } // Non-pointer receiver
func (dcf *DCFoo2) UnmarshalAmino(s string) { dcf.a = s }

func TestDeepCopyFoo2(t *testing.T) {
	dcf1 := newDCFoo2("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo2)
	assert.Equal(t, "foobar", dcf2.a)
}

type DCFoo3 struct{ a string }

func newDCFoo3(a string) *DCFoo3            { return &DCFoo3{a: a} }
func (dcf DCFoo3) MarshalAmino() string     { return dcf.a }
func (dcf *DCFoo3) UnmarshalAmino(s []byte) { dcf.a = string(s) } // Mismatch type

func TestDeepCopyFoo3(t *testing.T) {
	dcf1 := newDCFoo3("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo3)
	assert.Equal(t, "", dcf2.a)
}

type DCFoo4 struct{ a string }

func newDCFoo4(a string) *DCFoo4            { return &DCFoo4{a: a} }
func (dcf *DCFoo4) DeepCopy() *DCFoo4       { return &DCFoo4{"good"} }
func (dcf DCFoo4) MarshalAmino() string     { return dcf.a }
func (dcf *DCFoo4) UnmarshalAmino(s string) { dcf.a = string(s) } // Mismatch type

func TestDeepCopyFoo4(t *testing.T) {
	dcf1 := newDCFoo4("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo4)
	assert.Equal(t, "good", dcf2.a)
}

type DCFoo5 struct{ a string }

func newDCFoo5(a string) *DCFoo5            { return &DCFoo5{a: a} }
func (dcf DCFoo5) DeepCopy() DCFoo5         { return DCFoo5{"good"} }
func (dcf DCFoo5) MarshalAmino() string     { return dcf.a }
func (dcf *DCFoo5) UnmarshalAmino(s string) { dcf.a = string(s) } // Mismatch type

func TestDeepCopyFoo5(t *testing.T) {
	dcf1 := newDCFoo5("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo5)
	assert.Equal(t, "good", dcf2.a)
}

type DCFoo6 struct{ a string }

func newDCFoo6(a string) *DCFoo6     { return &DCFoo6{a: a} }
func (dcf *DCFoo6) DeepCopy() DCFoo6 { return DCFoo6{"good"} }

func TestDeepCopyFoo6(t *testing.T) {
	dcf1 := newDCFoo6("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo6)
	assert.Equal(t, "good", dcf2.a)
}

type DCFoo7 struct{ a string }

func newDCFoo7(a string) *DCFoo7     { return &DCFoo7{a: a} }
func (dcf DCFoo7) DeepCopy() *DCFoo7 { return &DCFoo7{"good"} }

func TestDeepCopyFoo7(t *testing.T) {
	dcf1 := newDCFoo7("foobar")
	dcf2 := amino.DeepCopy(dcf1).(*DCFoo7)
	assert.Equal(t, "good", dcf2.a)
}