[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-minimum-optional.svg)](https://pkg.go.dev/github.com/hymkor/go-minimum-optional)

go-minimal-optional
===================

This package has only two constructors (Some, None) and three methods (IfSome, IsNone, Match).

It requires Go 1.18 or later.

`optional.Option` is an array whose size is 0 or 1. Therefore, the contents can be handled with for-range even in versions below Go 1.22.

```example.go
package main

import (
    "github.com/hymkor/go-minimal-optional"
)

func test(x optional.Option[int]) {
    x.IfSome(func(v int) {
        println("   IfSome: it has a value:", v)
    })

    for _, v := range x {
        println("   for-range: it has a value:", v)
    }

    if x.IsNone() {
        println("   IsNone: it does not have a value")
    }

    x.Match(func(v int) {
        println("   Match: it has a value:", v)
    }, func() {
        println("   Match: it does not hava a value")
    })

    println()
}

func main() {
    println("None[int]")
    test(optional.None[int]())

    println("Some(4)")
    test(optional.Some(4))
}
```

**go run example.go**

```go run example.go|
None[int]
   IsNone: it does not have a value
   Match: it does not hava a value

Some(4)
   IfSome: it has a value: 4
   for-range: it has a value: 4
   Match: it has a value: 4

```
