package main

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for slow < len(nums) && fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
