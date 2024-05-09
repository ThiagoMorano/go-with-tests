package main

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lena")

	got := buffer.String()
	want := "Hello, Lena"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
