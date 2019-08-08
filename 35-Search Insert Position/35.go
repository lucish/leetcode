package thirty_five

func searchInsert(nums []int, target int) int {

	index, ok := binarySearch(nums, target)
	if ok {
		return index
	} else {
		if target > nums[index] {
			return index + 1
		} else {
			return index //- 1
		}
	}

}

func binarySearch(nums []int, target int) (index int, ok bool) {
	var left, right, mid int
	left, right = 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			//找到了  向前后分别找相同的
			return mid, true
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return mid, false
}
