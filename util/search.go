package util

func BinarySearch(nums []int, target int) int {

	var left, right, mid int
	left, right = 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			//找到了  向前后分别找相同的
			return mid
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1

}
