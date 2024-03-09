[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-minimal-optional.svg)](https://pkg.go.dev/github.com/hymkor/go-minimal-optional)

go-minimal-optional
===================

This package has only two constructors (Some, None) and three methods (IfSome, IsNone, Match).

It requires Go 1.18 or later.

+ `optional.Value` is an array whose size is 0 or 1. Therefore, the contents can be handled with for-range even in versions below Go 1.22.
+ `optional.Option` was renamed to `optional.Value`

```example.go
package main

import (
    "github.com/hymkor/go-minimal-optional"
)

func test(x optional.Value[int]) {
    x.IfSome(func(v int) {
        println("   IfSome: it has a value:", v)
    })

    for _, v := range x {
        println("   for-range(ready for v1.18): it has a value:", v)
    }

    for v := range x.Each {
        println("   for-range(v1.22 X:rangefunc): it has a value:", v)
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

```env GOEXPERIMENT=rangefunc go run example.go|
None[int]
   IsNone: it does not have a value
   Match: it does not hava a value

Some(4)
   IfSome: it has a value: 4
   for-range(ready for v1.18): it has a value: 4
   for-range(v1.22 X:rangefunc): it has a value: 4
   Match: it has a value: 4

```
