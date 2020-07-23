package palindromenumber

import (
	"testing"
)

type palindromeTestCase struct {
	x              int
	expectedOutput bool
}

func TestPalindrome(t *testing.T) {
	testCases := []palindromeTestCase{
		{-11, false},
		{0, true},
		{1, true},
		{10, false},
		{11, true},
		{123, false},
	}
	for i, testCase := range testCases {
		output := isPalindrome(testCase.x)
		if testCase.expectedOutput != output {
			t.Errorf("#%v: Expected %v, got %v; Input x = %v", i, testCase.expectedOutput, output, testCase.x)
		}
	}
}
