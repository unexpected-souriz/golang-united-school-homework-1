package main

import "testing"

func TestPrintEmoji(t *testing.T) {
	got := printEmoji()
	want := "Hello 🗺️ !"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
