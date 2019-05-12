package ulong

import "testing"

func TestNewInt(t *testing.T) {
	if 5 != int(NewInt(5)) {
		t.Error("not the same value, but should be")
	}
}

func TestNewInt64(t *testing.T) {
	if uint64(5) != uint64(NewInt64(5)) {
		t.Error("not the same value, but should be")
	}
}
