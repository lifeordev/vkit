package validy

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

func NotEmpty(value string) (*ValidationError, RuntimeError) {
	if value == "" {
		return NewValidationError(FailNotEmpty[0], FailNotEmpty[1]), nil
	}
	return nil, nil
}

func IsEmail(value string) (*ValidationError, RuntimeError) {
	_, err := mail.ParseAddress(value)
	// ParseAddress allows local domains, so we need to invalidate them additionally
	publicDomainRegex := regexp.MustCompile(`@[^@]*\.`)

	if err != nil || !publicDomainRegex.MatchString(value) {
		return NewValidationError(FailIsEmail[0], FailIsEmail[1]), nil
	}
	return nil, nil
}

func MinLength(length int) ValidationRule[string] {
	return func(value string) (*ValidationError, RuntimeError) {
		if len(value) < length {
			return NewValidationError(FailMinLength[0], fmt.Sprint(FailMinLength[1], length)), nil
		}
		return nil, nil
	}
}

func MaxLength(length int) ValidationRule[string] {
	return func(value string) (*ValidationError, RuntimeError) {
		if len(value) > length {
			return NewValidationError(FailMaxLength[0], fmt.Sprintf(FailMaxLength[1], length)), nil
		}
		return nil, nil
	}
}

func OneOf(values []string) ValidationRule[string] {
	return func(value string) (*ValidationError, RuntimeError) {
		for _, v := range values {
			if v == value {
				return nil, nil
			}
		}
		return NewValidationError(FailOneOf[0], fmt.Sprintf(FailOneOf[1], strings.Join(values, ", "))), nil
	}
}

func Regex(regex regexp.Regexp) ValidationRule[string] {
	return func(value string) (*ValidationError, RuntimeError) {
		if !regex.MatchString(value) {
			return NewValidationError(FailRegex[0], FailRegex[1]), nil
		}
		return nil, nil
	}
}

func EthAddress(value string) (*ValidationError, RuntimeError) {
	if !strings.HasPrefix(value, "0x") {
		return NewValidationError(FailEthAddress0x[0], FailEthAddress0x[1]), nil
	}
	if len(value) != 42 {
		return NewValidationError(FailEthAddressLength[0], FailEthAddressLength[1]), nil
	}
	// Ensure the address contains only valid hex characters
	isHex, err := regexp.MatchString("^(0x)[0-9a-fA-F]{40}$", value)
	if err != nil {
		return nil, err
	}

	if !isHex {
		return NewValidationError(FailEthHex[0], FailEthHex[1]), nil
	}

	return nil, nil
}

// Integer
func Min(min int) ValidationRule[int] {
	return func(value int) (*ValidationError, RuntimeError) {
		if value < min {
			return NewValidationError(FailMin[0], fmt.Sprintf(FailMin[1], min)), nil
		}
		return nil, nil
	}
}

func Max(max int) ValidationRule[int] {
	return func(value int) (*ValidationError, RuntimeError) {
		if value > max {
			return NewValidationError(FailMax[0], fmt.Sprintf(FailMax[1], max)), nil
		}
		return nil, nil
	}
}
