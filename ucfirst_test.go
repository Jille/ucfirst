package ucfirst

import "testing"

func TestUcfirst(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"hello world", "Hello world"},
		{"HELLO WORLD", "HELLO WORLD"},
		{"", ""},
	}
	for _, tc := range tests {
		got := Ucfirst(tc.in)
		if got != tc.want {
			t.Errorf("Ucfirst(%q): %q, want %q", tc.in, got, tc.want)
		}
	}
}
