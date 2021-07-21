package lengthOfLongestSubstring

import (
	"strings"
)

//Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(input_string string) int {
	length := len(input_string)
	if length <= 1 {
		return length
	}

	max_length, first_pointer, second_pointer, index, substring, target_char := 1, 0, 1, 0, "", ""

	for second_pointer < length {
		substring = input_string[first_pointer:second_pointer]
		target_char = input_string[second_pointer : second_pointer+1]
		index = strings.Index(substring, target_char)

		if index != -1 {
			first_pointer += index + 1
		}

		second_pointer++

		if max_length < second_pointer-first_pointer {
			max_length = second_pointer - first_pointer
		}
	}

	return max_length
}
