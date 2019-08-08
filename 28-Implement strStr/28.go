package twenty_eight

func strStr(haystack string, needle string) int {

	s1 := []byte(haystack)
	s2 := []byte(needle)
	l1, l2 := len(s1), len(s2)
	if l2 == 0 {
		return 0
	}

	//s1中寻找以s2开头相同的字符 并向后截取相同长度并进行比较
	tag := s2[0]
	var flag bool
	var i, j int
	for i = 0; i <= l1-l2; i++ {
		if s1[i] != tag {
			continue
		}
		//判断剩余byte是否相等
		flag = true
		for j = 1; j < l2; j++ {
			if s1[i+j] != s2[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
