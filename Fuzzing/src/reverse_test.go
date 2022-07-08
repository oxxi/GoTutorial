package main

import "testing"

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"hola, mundo", "odnum ,aloh"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Recerse: %q, want %q", rev, tc.want)
		}
	}

}
