func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1

	for left <= right {
		middle := left + ((right - left) / 2)

		value := nums[middle]
		if value == target {
			return middle
		}else if value > target {
			right = middle - 1
		}else {
			left = middle + 1
		}
	}
	
	return -1
}
