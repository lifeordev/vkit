package option

// Option represents an optional value.
type Option[T any] struct {
	value  T
	isSome bool
}

// Some creates an Option with a value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: value, isSome: true}
}

// None creates an Option without a value.
func None[T any]() Option[T] {
	var zero T
	return Option[T]{value: zero, isSome: false}
}

// IsSome returns true if the Option has a value.
func (o Option[T]) IsSome() bool {
	return o.isSome
}

// IsNone returns true if the Option does not have a value.
func (o Option[T]) IsNone() bool {
	return !o.isSome
}

// returns the value and a boolean indicating whether the value is present
func (o Option[T]) Value() (T, bool) {
	return o.value, o.isSome
}

// Unwrap returns the value if present, or panics if not.
func (o Option[T]) Unwrap() T {
	if !o.isSome {
		panic("called Unwrap on a None value")
	}
	return o.value
}

// UnwrapOr returns the value if present, or the provided default value if not.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if !o.isSome {
		return defaultValue
	}
	return o.value
}

// UnwrapOrElse returns the value if present, or computes it from the provided function if not.
func (o Option[T]) UnwrapOrElse(defaultFunc func() T) T {
	if !o.isSome {
		return defaultFunc()
	}
	return o.value
}
