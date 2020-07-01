package longestsubstring

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var maxSubstringLength, tail int
	tracker := make(map[rune]int)
	for head, r := range s {
		if loc, found := tracker[r]; found == false {
			substringLength := head - tail + 1
			if substringLength > maxSubstringLength {
				maxSubstringLength = substringLength
			}
		} else {
			for _, subRune := range s[tail:loc] {
				delete(tracker, subRune)
			}
			tail = loc + 1
		}
		tracker[r] = head
	}
	return maxSubstringLength
}
