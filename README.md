[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-minimum-optional.svg)](https://pkg.go.dev/github.com/hymkor/go-minimum-optional)

go-minimum-optional
===================

This package has two constructors and two methods only.

```main_test.go
package optional_test

import (
    "testing"

    "github.com/hymkor/go-minimum-optional"
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
```
