package dpi

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Renato")

	got := buffer.String()
	want := "Hello, Renato"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
