package thirty

func findSubstring(s string, words []string) []int {

	if len(words) == 0 {
		return []int{}
	}
	oneWordLen := len([]byte(words[0]))
	if oneWordLen == 0 {
		return []int{}
	}
	//每个单词长度相等 则匹配的子串长度应为单词个数*单个单词长度
	subLen := oneWordLen * len(words)
	src := []byte(s)
	srcLen := len(src)
	//初始化wordMap v为单词出现的次数
	wordMap := make(map[string]int, len(words))
	for _, elem := range words {
		v, has := wordMap[elem]
		if !has {
			wordMap[elem] = 1
		} else {
			wordMap[elem] = v + 1
		}
	}

	//将每个单词的个数存起来便于后续初始化
	initNum := make([]int, len(words))
	for i, elem := range words {
		initNum[i] = wordMap[elem]
	}

	var ret []int
	var flag bool
	for i := 0; i+subLen-1 < srcLen; i++ {
		flag = true
		buf := i
		//从i开始每隔单词长度进行一次截取 在wordmap中查找
		count := 0
		for {
			count++
			if count > len(words) {
				break
			}
			subByte := src[buf : buf+oneWordLen]
			subString := string(subByte)
			v, has := wordMap[subString]
			if !has {
				//单词不存在
				flag = false
				break
			}
			if v <= 0 {
				//此单词出现次数超过待匹配单词组中该单词的个数
				flag = false
				break
			} else {
				//此单词出现一次 个数减一
				wordMap[subString] = v - 1
			}
			//进行下一个单词匹配
			buf += oneWordLen
			// if buf+oneWordLen >= srcLen {
			// 	break
			// }
		}
		if flag == true {
			//匹配完成 是一个结果
			ret = append(ret, i)
		}

		//一次匹配后重置map
		for index, elem := range words {
			wordMap[elem] = initNum[index]
		}

	}

	return ret
}
