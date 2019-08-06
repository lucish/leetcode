package one

func twoSum(nums []int, target int) []int {

	numIndexMap := make(map[int]int, len(nums))

	for i, v := range nums {
		numIndexMap[v] = i
	}

	for i, v := range nums {
		antoherNum := target - v
		index, ok := numIndexMap[antoherNum]
		//另一个数存在且不重复
		if ok && i != index {
			array := []int{i, index}
			return array[:]
		}
	}

	return nil

}
