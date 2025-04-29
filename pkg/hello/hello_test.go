package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hello, World!"
	msg := SayHello()
	if msg != expected {
		t.Errorf(`SayHello() = %q, want %q`, msg, expected)
	}
}
