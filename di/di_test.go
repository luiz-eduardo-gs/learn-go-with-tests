package main

import (
	"bytes"
	"testing"
)

func TestGrret(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Luiz")

	got := buffer.String()
	want := "Hello, Luiz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
