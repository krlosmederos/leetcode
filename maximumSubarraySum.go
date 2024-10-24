package main

func maximumSubarraySum(nums []int, k int) int64 {
	l := 0
	ans, sum := int64(0), int64(0)
	seen := make(map[int]int)

	for r := 0; r < len(nums); r++ {
		seen[nums[r]]++
		sum += int64(nums[r])

		for seen[nums[r]] > 1 {
			seen[nums[l]]--
			sum -= int64(nums[l])
			l++
		}

		if r-l+1 == k {
			ans = max(ans, sum)
			sum -= int64(nums[l])
			seen[nums[l]]--
			l++
		}
	}

	return ans
}
