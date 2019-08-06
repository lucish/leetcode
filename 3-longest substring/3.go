package third

func lengthOfLongestSubstring(s string) int {

	//从最长的找
	src := []byte(s)
	len := len(src)
	if len == 0 {
		return 0
	}

	if len == 1 {
		return 1
	}

	//i为子串长度
	for i := len; i > 1; i-- {
		//子串为 [j, i+j)
		for j := 0; i+j <= len; j++ {
			des := src[j : i+j]
			if !hasRepeatByte(des) {
				return i
			}
		}

	}

	return 1
}

func hasRepeatByte(src []byte) bool {
	byteMap := make(map[byte]struct{}, len(src))
	for _, elem := range src {
		_, has := byteMap[elem]
		if has {
			return true
		}
		byteMap[elem] = struct{}{}
	}
	return false
}
