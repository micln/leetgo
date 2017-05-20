package main

func singleNonDuplicate(nums []int) int {
	size := len(nums) - 1
	for i := 0; i < size; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[size]
}

func singleNonDuplicate_Xor(nums []int) int {
	s := 0
	for i := range nums {
		s = s ^ nums[i]
	}
	return s
}
