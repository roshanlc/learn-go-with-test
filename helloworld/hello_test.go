package helloworld

import "testing"

func TestSayHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		name := "Ramu"
		got := sayHello(name)
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := sayHello("")
		want := "Hello, "

		assertCorrectMessage(t, got, want)
	})
}

// checks for assertion that got and want are equal
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q,want %q", got, want)
	}
}
