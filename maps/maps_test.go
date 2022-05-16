package maps

import "testing"

func TestSearch(t *testing.T) {
	dic := map[string]string{
		"test": "this is jst a test",
	}
	got := Search(dic, "test")
	want := "this is jst a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
