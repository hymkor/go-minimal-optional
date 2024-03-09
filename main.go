package optional

type Value[T any] []T

// Some is a constructor to make an instance that has a value.
func Some[T any](value T) Value[T] {
	return Value[T]([]T{value})
}

// None is a constructor to make an instance that does not have a value.
func None[T any]() Value[T] {
	return nil
}

// IfSome calls `f` when the receiver has a value. Otherwise it does nothing.
func (p Value[T]) IfSome(f func(value T)) {
	if len(p) >= 1 {
		f(p[0])
	}
}

func (p Value[T]) Each(f func(value T) bool) {
	if len(p) >= 1 {
		f(p[0])
	}
}

// IsNone returns true when the receiver does not has a value. Otherwise it returns false.
func (p Value[T]) IsNone() bool {
	return len(p) <= 0
}

func (p Value[T]) Match(then func(value T), otherwise func()) {
	if len(p) > 0 {
		then(p[0])
	} else {
		otherwise()
	}
}
