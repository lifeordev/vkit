package option

import (
	"encoding/json"
	"reflect"
)

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
// recommended way to extract value
func (o Option[T]) Unwrap() (T, bool) {
	return o.value, o.isSome
}

// Unwrap returns the value if present, or panics if not.
func (o Option[T]) ForceUnwrap() T {
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

func (o Option[_]) MarshalJSON() ([]byte, error) {
	if o.isSome {
		return json.Marshal(o.value)
	} else {
		return []byte("null"), nil
	}
}

func (o *Option[_]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.isSome = false

		return nil
	}

	err := json.Unmarshal(data, &o.value)

	if err != nil {
		return err
	}

	o.isSome = true

	return nil
}

// Checks if a given value is of type Option
func IsOption(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t != nil && t.Name() == "Option" {
		return true
	}
	return false
}

// Checks if a field inside a struct is of type Option
func IsFieldOfTypeOption(v interface{}, field string) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false
	}
	fieldVal := val.FieldByName(field)
	if !fieldVal.IsValid() {
		return false
	}

	fieldType := fieldVal.Type()
	return fieldType.Name() == "Option"
}
