package validy

type ValidationError struct {
	Code    string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(code string, message string) *ValidationError {
	return &ValidationError{
		Code:    code,
		Message: message,
	}
}

type invalid [2]string

var (
	FailNotEmpty         = invalid{"notEmpty.empty", "may not be empty."}
	FailIsEmail          = invalid{"isEmail.invalid", "must be a valid email address."}
	FailMinLength        = invalid{"minLength.length", "must be minimum %d characters long."}
	FailMaxLength        = invalid{"maxLength.length", "must be max %d characters long."}
	FailOneOf            = invalid{"oneOf.notFound", "is not one of %s."}
	FailRegex            = invalid{"regex.invalid", "does not match expected pattern."}
	FailEthAddress0x     = invalid{"eth.0x", "must start with '0x'."}
	FailEthAddressLength = invalid{"eth.length", "must be 42 characters long."}
	FailEthHex           = invalid{"eth.hex", "may only contain hexadecimal characters."}
	FailMin              = invalid{"min.invalid", "must be at least %d."}
	FailMax              = invalid{"max.invalid", "must not exceed %d."}
)
