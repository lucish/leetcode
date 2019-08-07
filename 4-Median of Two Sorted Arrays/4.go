package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)
	len := len1 + len2

	// fmt.Printf("len1: %v, len2: %v\n", len1, len2)

	// if len1 == 0 {

	// }
	// if len2 == 0 {

	// }

	var isEvenNum bool //总长度是否为偶数
	var index int      //中位数位置

	if len%2 == 0 {
		isEvenNum = true
		index = len / 2
	} else {
		isEvenNum = false
		index = (len + 1) / 2
	}

	if isEvenNum {
		//中位数为index 与 index+1 的平均数
		var isOnNums1 bool
		i, j := 0, 0
		var n1, n2 int
		for k := 1; k <= index; k++ {
			if i >= len1 {
				j++
				isOnNums1 = false
				continue
			}
			if j >= len2 {
				i++
				isOnNums1 = true
				continue
			}
			if nums1[i] <= nums2[j] {
				i++
				isOnNums1 = true
			} else {
				j++
				isOnNums1 = false
			}
			// fmt.Printf("i: %v, j: %v\n", i, j)
		}
		//找到inedx个是n1
		if isOnNums1 {
			n1 = nums1[i-1]
		} else {
			n1 = nums2[j-1]
		}
		//找到下一个是n2
		if i >= len1 {
			n2 = nums2[j]
			ret := (float64(n1) + float64(n2)) / 2
			return ret
		}
		if j >= len2 {
			n2 = nums1[i]
			ret := (float64(n1) + float64(n2)) / 2
			return ret
		}
		if nums1[i] <= nums2[j] {
			n2 = nums1[i]
		} else {
			n2 = nums2[j]
		}

		ret := (float64(n1) + float64(n2)) / 2
		return ret
	} else {
		//奇数个  找到index个即可
		var isOnNums1 bool
		i, j := 0, 0
		for k := 1; k <= index; k++ {
			if i >= len1 {
				j++
				isOnNums1 = false
				continue
			}
			if j >= len2 {
				i++
				isOnNums1 = true
				continue
			}
			if nums1[i] <= nums2[j] {
				i++
				isOnNums1 = true
			} else {
				j++
				isOnNums1 = false
			}
			// fmt.Printf("i: %v, j: %v\n", i, j)
		}
		if isOnNums1 {
			return float64(nums1[i-1])
		} else {
			return float64(nums2[j-1])
		}
	}

}

func main() {

	nums1 := make([]int, 2)
	nums2 := make([]int, 2)

	nums1[0], nums1[1] = 1, 2
	nums2[0], nums2[1] = 3, 4

	findMedianSortedArrays(nums1, nums2)
	return
}
