package validify

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

func NotEmpty(value string) (ValidationError, RuntimeError) {
	if value == "" {
		return errors.New("value cannot be empty"), nil
	}
	return nil, nil
}

func IsEmail(value string) (ValidationError, RuntimeError) {
	_, err := mail.ParseAddress(value)
	// ParseAddress allows local domains, so we need to invalidate them additionally
	publicDomainRegex := regexp.MustCompile(`@[^@]*\.`)

	if err != nil || !publicDomainRegex.MatchString(value) {
		return errors.New("invalid email format"), nil
	}
	return nil, nil
}

func MinLength(length int) ValidationRule[string] {
	return func(value string) (ValidationError, RuntimeError) {
		if len(value) < length {
			return fmt.Errorf("value must be at least %d characters long", length), nil
		}
		return nil, nil
	}
}

func MaxLength(length int) ValidationRule[string] {
	return func(value string) (ValidationError, RuntimeError) {
		if len(value) > length {
			return fmt.Errorf("value must be at most %d characters long", length), nil
		}
		return nil, nil
	}
}

func Min(min int) ValidationRule[int] {
	return func(value int) (ValidationError, RuntimeError) {
		if value < min {
			return fmt.Errorf("value must be at least %d", min), nil
		}
		return nil, nil
	}
}

func Max(max int) ValidationRule[int] {
	return func(value int) (ValidationError, RuntimeError) {
		if value > max {
			return fmt.Errorf("value must be at most %d", max), nil
		}
		return nil, nil
	}
}
