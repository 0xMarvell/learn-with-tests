package maps

import "testing"

func TestAdd(t *testing.T) {
	dic := Dictionary{"new": "adding a new word"}

	want := "adding a new word"
	got, err := dic.Add("new")
	if err != nil {
		t.Fatal("should be able to find the added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

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
