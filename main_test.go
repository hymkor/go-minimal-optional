package optional_test

import (
	"testing"

	"github.com/hymkor/go-minimal-optional"
)

func TestIfSome(t *testing.T) {
	called := false
	x := optional.Some(1)
	x.IfSome(func(v int) {
		if v != 1 {
			t.Fatal("value is not 1")
		}
		called = true
	})

	if !called {
		t.Fatal("IfSome is not called")
	}
}

func TestIsNone(t *testing.T) {
	x := optional.None[int]()

	x.IfSome(func(v int) {
		t.Fatal("IfSome is called for None")
	})

	if !x.IsNone() {
		t.Fatal("IsNone is false for None")
	}
}

func TestMatch(t *testing.T) {
	x := optional.Some(3)
	x.Match(func(v int) {
		if v != 3 {
			t.Fatal("Match: Some(3) must not call the first function with not 3")
		}
		//println("ok")
	}, func() {
		t.Fatal("Match: Some(3) must not call the second function")
	})

	y := optional.None[int]()
	y.Match(func(v int) {
		t.Fatal("Match: None[int] must not call the first function")
	}, func() {
		//println("ok")
	})
}
