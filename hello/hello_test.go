package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assert_correct_message := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %s  want: %s", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("NicX")
		want := "Hello, NicX"
		assert_correct_message(t, got, want)
	})

	t.Run("Say hello word, if name is empty string.", func(t *testing.T) {
		got := Hello("")
		want := "Hello, word"
		assert_correct_message(t, got, want)
	})
}
