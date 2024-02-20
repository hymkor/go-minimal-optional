//go:build ignore

package main

import (
	"github.com/hymkor/go-minimum-optional"
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
