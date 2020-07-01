package longestsubstring

import (
	"testing"
)

type testCase struct {
	input          string
	expectedOutput int
}

func TestLongesSubstring(t *testing.T) {
	testCases := []testCase{
		{"", 0},
		{"a", 1},
		{"ab", 2},
		{"abcab", 3},
		{"bbbbb", 1},
		{"abccbdacae", 4},
		{"pwwkew", 3},
		{"abcabcbb", 3},
		{"aaabcdec", 5},
	}
	for _, test := range testCases {
		output := lengthOfLongestSubstring(test.input)
		if output != test.expectedOutput {
			t.Errorf("Test case [%s]; Expected %v, got %v", test.input, test.expectedOutput, output)
		}
	}
}

func BenchmarkLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("aaabcdec")
	}
}
