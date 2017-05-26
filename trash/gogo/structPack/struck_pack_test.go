package structPack

import (
	"testing"
)

func TestStructPack(t *testing.T) {
	expStruct := NewExpStruct()
	t.Log(expStruct)
}

func TestNew(t *testing.T) {
	expStruct := New(1, 1)
	t.Log(expStruct)
	if expStruct.Mi1 != 1 {
		t.Fatal(expStruct)
	}
	if expStruct.Mf1 != 1 {
		t.Fatal(expStruct)
	}
}

const name = iota
const hello = iota

const (
	a = 1 << iota
	b
	c
	d
)

const (
	B float32  = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func TestIota(t *testing.T) {
	t.Log(name)
	t.Log(hello)
	t.Log(a)
	t.Log(b)
	t.Log(c)
	t.Log(d)
	t.Log(B, KB, MB, GB, TB, PB, EB, ZB, YB)
}
