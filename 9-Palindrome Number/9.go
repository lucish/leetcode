package ninth

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	num := make([]int, 0, 10)
	for {
		num = append(num, x%10)
		x /= 10
		if x == 0 {
			break
		}
	}

	len := len(num)
	i, j := 0, len-1
	for i+1 <= len/2 {
		if num[i] != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}
