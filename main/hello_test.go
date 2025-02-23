package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Zain", "")
		want := "Hello, Zain"
		assertCorrectMessage(t, got, want)
	})
	t.Run("when empty string is supplied - (Hello, World)", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Zain", "Spanish")
		var want string = "Hola, Zain"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Error got %q want %q", got, want)
	}
}
