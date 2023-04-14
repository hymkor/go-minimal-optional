package optional_test

import (
	"testing"

	"github.com/hymkor/go-minimum-optional"
)

func TestIfSome(t *testing.T) {
	found := false
	v := optional.New(1)
	v.IfSome(func(v int) {
		if v != 1 {
			t.Fatal("value is not 1")
		}
		found = true
	})

	if !found {
		t.Fatal("IfSome is not called")
	}
}

func TestIsNone(t *testing.T) {
	v := optional.None[int]()

	v.IfSome(func(v int) {
		t.Fatal("IfSome is called for None")
	})

	if !v.IsNone() {
		t.Fatal("IsNone is false for None")
	}
}
