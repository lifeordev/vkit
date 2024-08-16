package validy

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{
			name:    "preventLocal",
			args:    "abc@example",
			wantErr: true,
		}, {
			name:    "allowValid",
			args:    "abc@example.com",
			wantErr: false,
		}, {
			name:    "preventInvalid",
			args:    "whooops.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err, _ := IsEmail(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("IsEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEthAddress(t *testing.T) {

	tests := []struct {
		name     string
		arg      string
		wantVErr bool
	}{
		{
			name:     "validAddress",
			arg:      "0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
			wantVErr: false,
		},
		{
			name:     "emptyString",
			arg:      "",
			wantVErr: true,
		},
		{
			name:     "banana",
			arg:      "banana",
			wantVErr: true,
		},
		{
			name:     "0xbanana",
			arg:      "0xbanana",
			wantVErr: true,
		},
		{
			name:     "nonHex",
			arg:      "0xXXXXXXXXXXf860124dC4fEe278FDCBD3XXXXXXXX",
			wantVErr: true,
		},
		{
			name:     "wrongLength",
			arg:      "0x42d35cc663",
			wantVErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVErr, gotRErr := EthAddress(tt.arg)
			if tt.wantVErr && gotVErr == nil {
				t.Errorf("EthAddress(%s) expected ValidationError. Got nil", tt.arg)
			}
			if !tt.wantVErr && gotVErr != nil {
				t.Errorf("EthAddress(%s) expected NO ValidationError. Got %v", tt.arg, gotVErr)
			}
			if gotRErr != nil {
				t.Errorf("EthAddress(%s) returned RuntimeError: %v", tt.arg, gotRErr)
			}
		})
	}
}
