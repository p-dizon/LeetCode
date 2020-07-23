package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var reversed int
	for x > 0 {
		digit := x % 10
		reversed = (reversed * 10) + digit
		x = x / 10

		if reversed == x {
			// even number of digits
			return true
		} else if reversed == (x / 10) {
			// odd number of digits
			return true
		}
	}

	return false
}
