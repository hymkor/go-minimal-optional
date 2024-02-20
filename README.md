[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/go-minimum-optional.svg)](https://pkg.go.dev/github.com/hymkor/go-minimum-optional)

go-minimum-optional
===================

This package has two constructors and two methods only.

`optional.Option` is an array whose size is 0 or 1. Therefore, the contents can be handled with for-range even in versions below Go 1.22.

```example.go
package main

import (
    "github.com/hymkor/go-minimum-optional"
)

func test(x optional.Option[int]) {
    x.IfSome(func(v int) {
        println("   It has a value(callback):", v)
    })

    for _, v := range x {
        println("   It has a value(range):", v)
    }

    if x.IsNone() {
        println("   It does not have a value")
    }
    println()
}

func main() {
    println("None[int]")
    test(optional.None[int]())

    println("Some[int](4)")
    test(optional.Some(4))
}
```

**go run example.go**

```go run example.go|
None[int]
   It does not have a value

Some[int](4)
   It has a value(callback): 4
   It has a value(range): 4

```
