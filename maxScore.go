package main

func maxScore(cardPoints []int, k int) int {
	sum := 0
	ans := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	ans = max(ans, sum)

	for i := len(cardPoints) - 1; i >= 0 && k > 0; i-- {
		sum = (sum + cardPoints[i]) - (cardPoints[k-1])
		k--
		ans = max(ans, sum)
	}

	return ans
}
