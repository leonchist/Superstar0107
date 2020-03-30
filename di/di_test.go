package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Marcelo")

	got := buffer.String()
	want := "Hello, Marcelo"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
