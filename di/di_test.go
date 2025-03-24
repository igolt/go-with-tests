package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John Doe")

	got := buffer.String()
	expected := "Hello, John Doe"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
