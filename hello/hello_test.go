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
		got := Hello("NicX", "English")
		want := "Hello, NicX"
		assert_correct_message(t, got, want)
	})

	t.Run("Say hello word, if name is empty string.", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, word"
		assert_correct_message(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("許興宇", "Chinese")
		want := "哈摟, 許興宇"
		assert_correct_message(t, got, want)
	})
}
