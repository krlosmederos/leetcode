package main

func findSubstring(s string, words []string) []int {
	wordCounter := make(map[string]int)
	wordLen := len(words[0])
	var sol []int

	for _, word := range words {
		wordCounter[word]++
	}

	for i := 0; i < len(s)-len(words)*wordLen+1; i++ {
		seen := map[string]int{}
		for j := 0; j < len(words); j++ {
			idx := i + j*wordLen
			curWord := s[idx : idx+wordLen]

			if _, ok := wordCounter[curWord]; !ok {
				break
			}

			seen[curWord]++

			if seen[curWord] > wordCounter[curWord] {
				break
			}

			if j == len(words)-1 {
				sol = append(sol, i)
			}
		}
	}
	return sol
}
