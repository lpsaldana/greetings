package greetings

import (
	"regexp"
	"testing"
)

func TestHelloWithName(t *testing.T) {
	name := "TEST"

	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("TEST")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(
			`Hello("TEST") = %q, %v, want match for %#q, nil`,
			msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
