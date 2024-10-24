package main

func checkInclusion(s1 string, s2 string) bool {
	l, fixedLen := 0, len(s1)
	s1Freq, s2Freq := make(map[byte]int), make(map[byte]int)

	for _, letter := range s1 {
		s1Freq[byte(letter)]++
	}

	for r := 0; r < len(s2); r++ {
		s2Freq[s2[r]]++

		if r-l+1 == fixedLen {
			isFound := true
			for ch := 'a'; ch <= 'z'; ch++ {
				if s1Freq[byte(ch)] != s2Freq[byte(ch)] {
					isFound = false
					break
				}
			}

			if isFound {
				return true
			}

			s2Freq[s2[l]]--
			l++
		}

	}
	return false
}
