package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Yash")
	want := "Hello, Yash!"
	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
