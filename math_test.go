package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", total, 30)
	}
}
