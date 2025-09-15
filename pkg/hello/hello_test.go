package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hello, World! Have a great day!"
	msg := SayHello("Have a great day!")
	if msg != expected {
		t.Errorf(`SayHello() = %q, want %q`, msg, expected)
	}
}
