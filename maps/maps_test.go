package maps

import (
	"testing"
)

func TestSearchDictionary(t *testing.T) {

	dict := Dictionary{"hi": "a form of greeting"}
	word := "hi"

	t.Run("known word", func(t *testing.T) {
		got := dict.SearchDictionary(word)
		want := "a form of greeting"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got := dict.SearchDictionary("hey")
		want := "yet another form of greeting"

		assertStrings(t, got, want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
