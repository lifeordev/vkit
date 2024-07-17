package validify

import (
	"testing"

	"github.com/lifeordev/vkit/option"
)

func TestValidateField(t *testing.T) {
	tests := []struct {
		name       string
		result     FieldValidationResult
		nrOfVErr   int
		expectRErr bool
	}{
		{
			name:       "ValidString",
			result:     ValidateField("foo", "Foo", MinLength(2)),
			nrOfVErr:   0,
			expectRErr: false,
		}, {
			name:       "InvalidString",
			result:     ValidateField("bar", "Bar", MaxLength(2), MinLength(50)),
			nrOfVErr:   2,
			expectRErr: false,
		}, {
			name:       "ValidInt",
			result:     ValidateField("foo", 33, Max(99)),
			nrOfVErr:   0,
			expectRErr: false,
		}, {
			name:       "InvalidInt",
			result:     ValidateField("foo", 33, Max(2), Min(1)),
			nrOfVErr:   1,
			expectRErr: false,
		}, {
			name:       "ValidOptionSome",
			result:     ValidateOptionField("foo", option.Some("test"), MinLength(2)),
			nrOfVErr:   0,
			expectRErr: false,
		}, {
			name:       "ValidOptionNone",
			result:     ValidateOptionField("foo", option.None[int](), Max(3)),
			nrOfVErr:   0,
			expectRErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectRErr && tt.result.RuntimeError != nil {
				t.Errorf("TestValidateField()/%s: Received unexpected RuntimeError (%s)",
					tt.name,
					tt.result.RuntimeError.Error(),
				)
			}
			if tt.expectRErr && tt.result.RuntimeError == nil {
				t.Errorf("TestValidateField()/%s: Expected RuntimeError, received nil", tt.name)
			}
			if tt.nrOfVErr != len(tt.result.ValidationErrors) {
				t.Errorf("TestValidateField()/%s: Unexpected number of ValidationErrors (expected: %d, received: %d)",
					tt.name,
					tt.nrOfVErr,
					len(tt.result.ValidationErrors),
				)
			}
		})
	}
}
