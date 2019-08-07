package eleven

func maxArea(height []int) int {

	l := len(height)
	res := 0
	i, j := 0, l-1
	var buf int
	for i < j {
		if height[i] < height[j] {
			buf = height[i] * (j - i)
			if buf > res {
				res = buf
			}
			i++
		} else {
			buf = height[j] * (j - i)
			if buf > res {
				res = buf
			}
			j--
		}
	}
	return res

}
