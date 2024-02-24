package main;

// O(n)
// Building hashmap while walking through array.
func twoSumN(nums []int, target int) []int {
    hashMap := make(map[int]int, len(nums))

    for i,v := range(nums){
        t := target - v
        value, exists := hashMap[t]
        if (exists){
            return []int{value, i}
        }
        hashMap[v] = i
    }
    return []int {0,0}
}

// O(n^2)
func twoSumNSquare(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
