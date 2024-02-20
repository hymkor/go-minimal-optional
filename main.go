package optional

type Option[T any] []T

// Some is a constructor to make an instance that has a value.
func Some[T any](value T) Option[T] {
	return Option[T]([]T{value})
}

// None is a constructor to make an instance that does not have a value.
func None[T any]() Option[T] {
	return nil
}

// IfSome calls `f` when the receiver has a value. Otherwise it does nothing.
func (p Option[T]) IfSome(f func(value T)) {
	if len(p) >= 1 {
		f(p[0])
	}
}

// IsNone returns true when the receiver does not has a value. Otherwise it returns false.
func (p Option[T]) IsNone() bool {
	return len(p) <= 0
}
