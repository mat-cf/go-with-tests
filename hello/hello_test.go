package main

import "testing"

// testing.TB is a interface satisfied by testing.T and testing.B . 
// When you have more than one argument of the same type (in our case two strings) 
// rather than having (got string, want string) you can shorten it to (got, want string).
func assertCorrectMessage(t testing.TB, got, want string) {
	// ensures that when it fails, it points to the function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Maria", "Portuguese")
		want := "Olá, Maria"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
