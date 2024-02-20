//go:build ignore

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
