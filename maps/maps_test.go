package maps

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is jst a test"}
	got := dic.Search("test")
	want := "this is jst a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
