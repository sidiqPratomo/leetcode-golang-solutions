package longestpalindrome

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := expandFromCenter(s, i, i)
		s2 := expandFromCenter(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func expandFromCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}