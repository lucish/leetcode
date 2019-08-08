package thirty_four

func searchRange(nums []int, target int) []int {

	var from, to int

	var left, right, mid int

	left, right = 0, len(nums)-1

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			//找到了  向前后分别找相同的
			from, to = mid, mid
			for from >= 1 && nums[from-1] == target {
				from--
			}
			for to+1 < len(nums) && nums[to+1] == target {
				to++
			}
			return []int{from, to}
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return []int{-1, -1}
}
