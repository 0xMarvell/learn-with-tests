package maps

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is jst a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is jst a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dic.Search("unknown")
		assertErr(t, got, ErrNotFound)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErr(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
