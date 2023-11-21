package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("badValue")
		assertError(t, got, ErrNotFound)
	})
	t.Run("add word", func(t *testing.T) {
		err := dict.addWord("newWord", "this is a new word")
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		got, err := dict.Search("newWord")
		want := "this is a new word"
		if err != nil {
			t.Errorf("should find word: %s", err)
		}
		assertString(t, got, want)
	})
	t.Run("existing word", func(t *testing.T) {
		err := dict.addWord("test", "this is a new word")
		assertError(t, err, ErrKeyPresent)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
