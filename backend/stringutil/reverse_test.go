package stringutil

import "testing"

func TEST001(t *testing.T) {
	input := "Hello"
	output := "olleH"

	if output != Reverse(input) {
		t.Errorf("Reversal test for input : %q failed with output %q", input, Reverse(input))
	}
}