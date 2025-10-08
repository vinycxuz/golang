package challenges

/*
Statement
Write a function that takes a string as input and checks whether it can be a valid palindrome by removing at most one character from it.
*/

func isPalindrome(str string) bool {
	isPalindromeIn := func(left, right int) bool {
		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return isPalindromeIn(left+1, right) || isPalindromeIn(left, right-1)
		}
		left++
		right--
	}
	return true
}
