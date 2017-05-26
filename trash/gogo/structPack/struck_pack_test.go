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
