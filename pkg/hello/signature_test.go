package hello

import (
	"testing"
)

func TestSignatureVariable(t *testing.T) {
	var myFuncVariable func(string) int

	myFuncVariable = f1
	result := myFuncVariable("Hello")
	if result != 5 {
		t.Errorf("f1: got %d, want 5", result)
	}

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	if result != 500 {
		t.Errorf("f2: got %d, want 500", result)
	}
}
