package validy

import "github.com/lifeordev/vkit/option"

type RuntimeError error
type ValidationRule[T any] func(value T) (*ValidationError, RuntimeError)

// ValidationAggregate aggregates validation errors for multiple fields.
type ValidationAggregate struct {
	ValidationErrors map[string][]ValidationError
}

// FieldValidationResult represents the result of validating a single field.
type FieldValidationResult struct {
	Field            string
	ValidationErrors []ValidationError
	RuntimeError     RuntimeError
}

// AggregateFieldValidation aggregates multiple FieldValidationResult instances into a single ValidationAggregate.
// If any FieldValidationResult contains a RuntimeError, the function will return immediately with that RuntimeError.
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

// ValidateField validates a single field value against a set of ValidationRules.
// Returns a FieldValidationResult containing any validation errors and a RuntimeError if one occurred.
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
			result.ValidationErrors = append(result.ValidationErrors, *vErr)
		}
	}
	return result
}

// ValidateOptionField validates a field value wrapped in an option.Option against a set of ValidationRules.
// If the option is empty, it returns an empty FieldValidationResult.
// Returns a FieldValidationResult containing any validation errors and a RuntimeError if one occurred.
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
			result.ValidationErrors = append(result.ValidationErrors, *vErr)
		}
	}
	return result
}
