package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Allen"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Allen")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Allen") = %q, %v, want match for %#q, nill`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
