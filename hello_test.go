package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Marcelo")
	want := "Hello, Marcelo"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}