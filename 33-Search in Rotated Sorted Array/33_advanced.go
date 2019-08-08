package thirty_three

func search1(nums []int, target int) int {

	var left, right, mid int

	if len(nums) == 0 {
		return -1
	}

	//首选寻找旋转的地方
	left, right = 0, len(nums)-1

	var rotatedIndex int

	if len(nums) >= 2 && nums[left] >= nums[right] {
		for left <= right {

			mid = (left + right) / 2
			if nums[mid] > nums[mid+1] {
				rotatedIndex = mid + 1
				break
			} else {
				if nums[mid] >= nums[left] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}

	// fmt.Println("rotatedIndex: ", rotatedIndex)

	if target >= nums[0] {
		if rotatedIndex != 0 {
			left, right = 0, rotatedIndex
		} else {
			left, right = 0, len(nums)-1
		}
	} else {
		left, right = rotatedIndex, len(nums)-1
	}

	for left <= right {
		mid = (left + right) / 2
		// fmt.Printf("%v %v %v\n", left, right, mid)
		if nums[mid] == target {
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

// func main() {
// 	nums := []int{1, 3}
// 	fmt.Println(search(nums, 3))
// }
