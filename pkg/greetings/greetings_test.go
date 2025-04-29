package greetings

import (
	"regexp"
	"testing"
)

// TestHiName calls greetings.Hi with a name, checking
// for a valid return value.
func TestHiName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hi("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hi("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHiEmpty calls greetings.Hi with an empty string,
// checking for an error.
func TestHiEmpty(t *testing.T) {
	msg, err := Hi("")
	if msg != "" || err == nil {
		t.Errorf(`Hi("") = %q, %v, want "", error`, msg, err)
	}
}
