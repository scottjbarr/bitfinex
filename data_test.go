package bitfinex

import "testing"

func TestStripStringData(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "mixed",
			input: "[\"foo",
			want:  "foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripStringData(tt.input)

			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
