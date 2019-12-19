package gomodversion

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "**HELLO**"
	got := Hello()

	if expected != got {
		t.Errorf("expected: %v\tgot: %v", expected, got)
	}
}
