package main

import "fmt"

func totalFruit(fruits []int) int {
	ans, l := 1, 0
	seen := make(map[int]int)

	for r := 0; r < len(fruits); r++ {
		seen[fruits[r]]++
		fmt.Println(len(seen))
		for len(seen) > 2 {
			seen[fruits[l]]--
			if seen[fruits[l]] == 0 {
				delete(seen, fruits[l])
			}
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}
