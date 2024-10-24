package main

func minWindow(s string, t string) string {
	m := make(map[byte]int)
	win_l := -1
	l, seen := 0, len(t)
	win_len := 1000001

	for i := 0; i < len(t); i++ {
		m[t[i]] = m[t[i]] + 1
	}

	for r := 0; r < len(s); r++ {
		m[s[r]]--
		if v, ok := m[s[r]]; ok {
			if v >= 0 {
				seen--
			}
		}

		for seen == 0 {
			if r-l+1 < win_len {
				win_l = l
				win_len = r - l + 1
			}
			m[s[l]]++
			if m[s[l]] > 0 {
				seen++
			}
			l++
		}
	}

	if win_l == -1 {
		return ""
	}

	return string(s[win_l : win_l+win_len])
}
