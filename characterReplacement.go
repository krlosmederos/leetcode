package main

func characterReplacement(s string, k int) int {
	ans, l := 0, 0
	seen := make(map[byte]int)
	maxFreqLetter := 0

	for r := 0; r < len(s); r++ {
		seen[s[r]]++
		if maxFreqLetter < seen[s[r]] {
			maxFreqLetter = seen[s[r]]
		}

		if r+1-l-maxFreqLetter > k {
			seen[s[l]]--
			l++
		}

		if ans < r-l+1 {
			ans = r - l + 1
		}
	}

	return ans
}
