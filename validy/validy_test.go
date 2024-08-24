package validy

import (
	"testing"

	"github.com/lifeordev/vkit/option"
)

func TestValidateField(t *testing.T) {
	tests := []struct {
		name       string
		result     FieldValidationResult
		expectVErr bool
		expectRErr bool
	}{
		{
			name:       "ValidString",
			result:     ValidateField("foo", "Foo", MinLength(2)),
			expectVErr: false,
			expectRErr: false,
		}, {
			name:       "InvalidString",
			result:     ValidateField("bar", "Bar", MaxLength(2), MinLength(50)),
			expectVErr: true,
			expectRErr: false,
		}, {
			name:       "ValidInt",
			result:     ValidateField("foo", 33, Max(99)),
			expectVErr: false,
			expectRErr: false,
		}, {
			name:       "InvalidInt",
			result:     ValidateField("foo", 33, Max(2), Min(1)),
			expectVErr: true,
			expectRErr: false,
		}, {
			name:       "ValidOptionSome",
			result:     ValidateOptionField("foo", option.Some("test"), MinLength(2)),
			expectVErr: false,
			expectRErr: false,
		}, {
			name:       "ValidOptionNone",
			result:     ValidateOptionField("foo", option.None[int](), Max(3)),
			expectVErr: false,
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
			if tt.expectVErr && tt.result.ValidationError == nil {
				t.Errorf("TestValidateField()/%s: Expected ValidationError, received nil", tt.name)
			}
		})
	}
}
