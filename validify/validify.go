package validify

import "github.com/lifeordev/vkit/option"

// / Actual Library
type ValidationError error
type RuntimeError error
type ValidationRule[T any] func(value T) (ValidationError, RuntimeError)

type ValidationAggregate struct {
	ValidationErrors map[string][]ValidationError
}

type FieldValidationResult struct {
	Field            string
	ValidationErrors []ValidationError
	RuntimeError     RuntimeError
}

func AggregateFieldValidation(results ...FieldValidationResult) (ValidationAggregate, RuntimeError) {
	aggregate := ValidationAggregate{
		ValidationErrors: make(map[string][]ValidationError),
	}
	for _, fieldResult := range results {
		if fieldResult.RuntimeError != nil {
			return aggregate, fieldResult.RuntimeError
		}
		aggregate.ValidationErrors[fieldResult.Field] = fieldResult.ValidationErrors
	}
	return aggregate, nil
}

func ValidateField[T any](field string, value T, rules ...ValidationRule[T]) FieldValidationResult {
	result := FieldValidationResult{
		Field:            field,
		ValidationErrors: make([]ValidationError, 0),
		RuntimeError:     nil,
	}

	for _, rule := range rules {
		vErr, err := rule(value)
		if err != nil {
			// Exit on first Runtime Error, no need to execute any other validators
			result.RuntimeError = err
			return result
		}
		if vErr != nil {
			result.ValidationErrors = append(result.ValidationErrors, vErr)
		}
	}
	return result
}

func ValidateOptionField[T any](field string, optionValue option.Option[T], rules ...ValidationRule[T]) FieldValidationResult {
	result := FieldValidationResult{
		Field:            field,
		ValidationErrors: make([]ValidationError, 0),
		RuntimeError:     nil,
	}
	value, ok := optionValue.Unwrap()
	if !ok {
		return result
	}

	for _, rule := range rules {
		vErr, err := rule(value)
		if err != nil {
			// Exit on first Runtime Error, no need to execute any other validators
			result.RuntimeError = err
			return result
		}
		if vErr != nil {
			result.ValidationErrors = append(result.ValidationErrors, vErr)
		}
	}
	return result
}
