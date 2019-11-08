package main

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicate(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return len(nums)
}

func main() {

}
