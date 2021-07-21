package lengthOfLongestSubstring

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert_correct_message := func(t *testing.T, result int, expect int) {
		t.Helper()
		if result != expect {
			t.Errorf("result: %d, expect: %d", result, expect)
		}
	}
	t.Run("Test case 1: Input is abcabcbb", func(t *testing.T) {
		result := LengthOfLongestSubstring("abcabcbb")
		expect := 3
		assert_correct_message(t, result, expect)
	})

	t.Run("Test case 2: Input is bbbbb", func(t *testing.T) {
		result := LengthOfLongestSubstring("bbbbb")
		expect := 1
		assert_correct_message(t, result, expect)
	})

	t.Run("Test case 3: Input is pwwkew", func(t *testing.T) {
		result := LengthOfLongestSubstring("pwwkew")
		expect := 3
		assert_correct_message(t, result, expect)
	})

	t.Run("Test case 4: Input is empty", func(t *testing.T) {
		result := LengthOfLongestSubstring("")
		expect := 0
		assert_correct_message(t, result, expect)
	})
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstring("abcabcbb")
	}
}

func ExampleLengthOfLongestSubstring() {
	result := LengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
	//Output: 3
}
