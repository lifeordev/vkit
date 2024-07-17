package validify

import "testing"

func TestIsEmail(t *testing.T) {
	type args struct {
		value string
	}
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
