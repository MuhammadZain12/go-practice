package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"path": "This is our path to valhalla"}
	got := "This is our path to valhalla"
	want := Search(dictionary, "path")

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
