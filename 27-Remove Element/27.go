package twenty_seven

func removeElement(nums []int, val int) int {

	var i, j int

	l := len(nums)

	for j < l {
		if nums[j] != val {
			nums[i] = nums[j]
			j++
			i++
		} else {
			j++
		}

	}
	return i
}
