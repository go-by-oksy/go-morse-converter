package service

import (
	"errors"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:  "text to Morse",
			input: "Привет",
			want:  ".--. .-. .. .-- . -",
		},
		{
			name:  "Morse to text",
			input: ".--. .-. .. .-- . -",
			want:  "ПРИВЕТ",
		},
		{
			name:  "text containing a period",
			input: "Привет.",
			want:  ".--. .-. .. .-- . - ......",
		},
		{
			name:    "empty input",
			input:   "   ",
			wantErr: ErrEmptyData,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Convert(test.input)

			if test.wantErr != nil {
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("Convert() error = %v, want %v", err, test.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Convert() unexpected error: %v", err)
			}

			if got != test.want {
				t.Errorf("Convert() = %q, want %q", got, test.want)
			}
		})
	}
}
