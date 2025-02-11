package main

import "testing"

func TestAdder(t *testing.T) {
	total := adder(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
