package longestsubstringwithoutrepeatingchar

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	hash := make(map[byte]int)
	max := 0
	for i, j := 0, -1; i < len(s); i++ {
		if val, ok := hash[s[i]]; ok {
			j = int(math.Max(float64(j), float64(val)))
		}
		hash[s[i]] = i
		max = int(math.Max(float64(max), float64(i-j)))
	}
	return max
}