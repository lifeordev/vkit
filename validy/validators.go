package validy

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

type CompatibleTypes interface {
	string | int
}

// When applies validation rules if the 'when' function returns true for the given value.
func When[T CompatibleTypes](when func(T) bool, rules ...ValidationRule[T]) ValidationRule[T] {
	return func(value T) (*ValidationError, *RuntimeError) {
		if when(value) {
			for _, rule := range rules {
				vErr, rErr := rule(value)
				if vErr != nil || rErr != nil {
					return vErr, rErr
				}
			}
		}
		return nil, nil
	}
}

// WhenNotEmpty applies the given validation rules only when the provided value is not empty.
// This function is generic and works with types that conform to the CompatibleTypes constraint.
//
// The function considers a value to be "empty" based on the following criteria:
//   - For strings, a value is empty if it is an empty string ("").
//   - For integers, a value is empty if it is zero (0).
//   - For other types, the value is considered non-empty by default.
func WhenNotEmpty[T CompatibleTypes](rules ...ValidationRule[T]) ValidationRule[T] {
	when := func(value T) bool {
		switch v := any(value).(type) {
		case string:
			return v == ""
		case int:
			return v == 0
		}
		return true
	}
	return When(when, rules...)
}

// #### String Validators

// NotEmpty validates that a string is not empty. Returns a ValidationError if the string is empty, otherwise nil.
// Example: NotEmpty("") returns an error indicating the string cannot be empty.
func NotEmpty(value string) (*ValidationError, *RuntimeError) {
	if value == "" {
		return NewValidationError(FailNotEmpty[0], FailNotEmpty[1]), nil
	}
	return nil, nil
}

// IsEmail validates whether the given string is a valid email address.
func IsEmail(value string) (*ValidationError, *RuntimeError) {
	_, err := mail.ParseAddress(value)
	// ParseAddress allows local domains, so we need to invalidate them additionally
	publicDomainRegex := regexp.MustCompile(`@[^@]*\.`)

	if err != nil || !publicDomainRegex.MatchString(value) {
		return NewValidationError(FailIsEmail[0], FailIsEmail[1]), nil
	}
	return nil, nil
}

// MinLength returns a validation rule that ensures a string is at least the specified length.
// Returns a ValidationError if the string is shorter than the given length, otherwise nil.
// Example: MinLength(5) checks if the string is at least 5 characters long.
func MinLength(length int) ValidationRule[string] {
	return func(value string) (*ValidationError, *RuntimeError) {
		if len(value) < length {
			return NewValidationError(FailMinLength[0], fmt.Sprint(FailMinLength[1], length)), nil
		}
		return nil, nil
	}
}

// MaxLength returns a validation rule that ensures a string does not exceed the specified length.
// Returns a ValidationError if the string is longer than the given length, otherwise nil.
// Example: MaxLength(10) checks if the string is at most 10 characters long.
func MaxLength(length int) ValidationRule[string] {
	return func(value string) (*ValidationError, *RuntimeError) {
		if len(value) > length {
			return NewValidationError(FailMaxLength[0], fmt.Sprintf(FailMaxLength[1], length)), nil
		}
		return nil, nil
	}
}

// OneOf returns a validation rule that ensures a string matches one of the specified values.
// Returns a ValidationError if the string does not match any of the given values, otherwise nil.
// Example: OneOf([]string{"red", "blue"}) checks if the string is either "red" or "blue".
func OneOf(values []string) ValidationRule[string] {
	return func(value string) (*ValidationError, *RuntimeError) {
		for _, v := range values {
			if v == value {
				return nil, nil
			}
		}
		return NewValidationError(FailOneOf[0], fmt.Sprintf(FailOneOf[1], strings.Join(values, ", "))), nil
	}
}

// Regex returns a validation rule that ensures a string matches the specified regular expression.
// Returns a ValidationError if the string does not match the regex, otherwise nil.
// Example: Regex(regexp.MustCompile(`^\d+$`)) checks if the string contains only digits.
func Regex(regex regexp.Regexp) ValidationRule[string] {
	return func(value string) (*ValidationError, *RuntimeError) {
		if !regex.MatchString(value) {
			return NewValidationError(FailRegex[0], FailRegex[1]), nil
		}
		return nil, nil
	}
}

// EthAddress validates an Ethereum address by checking if it starts with "0x", has a length of 42 characters,
// and contains only valid hex characters after the "0x" prefix. Returns a ValidationError if any check fails, otherwise nil.
func EthAddress(value string) (*ValidationError, *RuntimeError) {
	if !strings.HasPrefix(value, "0x") {
		return NewValidationError(FailEthAddress0x[0], FailEthAddress0x[1]), nil
	}
	if len(value) != 42 {
		return NewValidationError(FailEthAddressLength[0], FailEthAddressLength[1]), nil
	}
	// Ensure the address contains only valid hex characters
	isHex, err := regexp.MatchString("^(0x)[0-9a-fA-F]{40}$", value)
	if err != nil {
		return nil, NewRuntimeError(err.Error())
	}

	if !isHex {
		return NewValidationError(FailEthHex[0], FailEthHex[1]), nil
	}

	return nil, nil
}

// ####### Integer Validators

// Min returns a validation rule that ensures an integer is at least the specified minimum value.
// Returns a ValidationError if the integer is less than the minimum, otherwise nil.
// Example: Min(10) checks if the integer is at least 10.
func Min(min int) ValidationRule[int] {
	return func(value int) (*ValidationError, *RuntimeError) {
		if value < min {
			return NewValidationError(FailMin[0], fmt.Sprintf(FailMin[1], min)), nil
		}
		return nil, nil
	}
}

// Max returns a validation rule that ensures an integer does not exceed the specified maximum value.
// Returns a ValidationError if the integer is greater than the maximum, otherwise nil.
// Example: Max(100) checks if the integer is at most 100.
func Max(max int) ValidationRule[int] {
	return func(value int) (*ValidationError, *RuntimeError) {
		if value > max {
			return NewValidationError(FailMax[0], fmt.Sprintf(FailMax[1], max)), nil
		}
		return nil, nil
	}
}
