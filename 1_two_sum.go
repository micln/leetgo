package main

func twoSum(nums []int, target int) []int {

	pos := make(map[int]int, len(nums))

	for i := range nums {
		diff := target - nums[i]
		if pos[diff] > 0 {
			return []int{pos[diff] - 1, i}
		}

		pos[nums[i]] = i + 1
	}

	return nil
}
