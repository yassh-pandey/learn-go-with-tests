package main

import "testing"

func TestGreet(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q want: %q", got, want)
		}
	}
	t.Run("Greeting with argument", func(t *testing.T) {
		got := Greet("Yash", "english")
		want := "Hello, Yash!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Greeting with empty string", func(t *testing.T) {
		got := Greet("", "english")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Greet in Spanish", func(t *testing.T) {
		got := Greet("Alba", "spanish")
		want := "Hola, Alba!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Greet in French", func(t *testing.T) {
		got := Greet("Amelia", "french")
		want := "Bonjour, Amelia!"
		assertCorrectMessage(t, got, want)
	})
}
