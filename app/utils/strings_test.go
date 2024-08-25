package strings

import (
	"testing"
)

var testText = "test"

func TestEmptyToNil(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "Empty string",
			args: args{
				str: "",
			},
			want: nil,
		},
		{
			name: "Non-empty string",
			args: args{
				str: testText,
			},
			want: &testText,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EmptyToNil(tt.args.str)
			if got == nil && tt.want == nil {
				return
			}
			if *got != *tt.want {
				t.Errorf("EmptyToNil() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestIsValidHexColor(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsValid color",
			args: args{
				color: "#ffffff",
			},
			want: false,
		},
		{
			name: "IsInvalid color",
			args: args{
				color: "#00000",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidHexColor(tt.args.color); got != tt.want {
				t.Errorf("IsValidHexColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
