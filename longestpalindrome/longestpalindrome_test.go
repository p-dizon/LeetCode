package longestpalindrome

import "testing"

type isPalindromeTestCase struct {
	input          string
	expectedOutput bool
}

type longestPalindromeTestCase struct {
	input          string
	expectedOutput string
}

func TestIsPalindrome(t *testing.T) {
	testCases := []isPalindromeTestCase{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abc", false},
		{"aabbaa", true},
	}
	for _, test := range testCases {
		output := isPalindrome(test.input)
		if output != test.expectedOutput {
			t.Errorf("Test case [%s]; Expected %v, got %v", test.input, test.expectedOutput, output)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	testCases := []longestPalindromeTestCase{
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"aba", "aba"},
		{"abc", "a"},
		{"aabbaa", "aabbaa"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"abcdeefeeg", "eefee"},
	}
	for _, test := range testCases {
		output := longestPalindrome(test.input)
		if output != test.expectedOutput {
			t.Errorf("Test case [%s]; Expected %v, got %v", test.input, test.expectedOutput, output)
		}
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome("abcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeegabcdeefeeg")
	}
}
