package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("use the prefix 'Hola, ' when the language is 'Spanish'", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("use the prefix 'Bonjour, ' when the language is 'French'", func(t *testing.T) {
		got := Hello("François", "French")
		want := "Bonjour, François!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
