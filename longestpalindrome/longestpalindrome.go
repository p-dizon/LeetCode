package longestpalindrome

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	var longestSubstring string
	if len(s) == 1 {
		return s
	}
	for i := range s {
		for j := len(s); j > i; j-- {
			if s[j-1] != s[i] {
				continue
			} else if false == isPalindrome(s[i:j]) {
				continue
			}
			if j-i > len(longestSubstring) {
				longestSubstring = s[i:j]
			}
		}
	}
	return longestSubstring
}
