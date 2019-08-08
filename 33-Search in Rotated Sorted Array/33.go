package thirty_three

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	// head := nums[0]

	var isTargetOnLeft bool
	if target >= nums[0] {
		isTargetOnLeft = true
	} else {
		isTargetOnLeft = false
	}

	var start, end, mid int
	start, end = 0, len(nums)-1
	for start <= end {
		mid = (start + end) / 2
		// fmt.Println("start: ", start, " end: ", end)
		// fmt.Printf("%v: %v\n", mid, nums[mid])
		if nums[mid] == target {
			return mid
		}

		//在左区间或者右区间 均为升序 直接二分查找
		if nums[start] <= nums[end] {
			if target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { //横跨两个区间 一步一步削减另一半区间
			if isTargetOnLeft {
				end--
			} else {
				start++
			}
		}

	}

	return -1
}
