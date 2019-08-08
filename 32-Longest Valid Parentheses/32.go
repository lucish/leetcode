package thirty_two

type node struct {
	Elem  byte
	Index int
}

//使用数组模拟一个栈  top指向栈顶
var top int
var stack []*node
var indexArr []bool

func longestValidParentheses(s string) int {

	src := []byte(s)
	//下表数组 值表示是否包括在匹配的括号之中(包括括号)
	indexArr = make([]bool, len(src))

	//栈初始化
	initStack(len(src))

	//进行括号匹配
	for i := 0; i < len(src); i++ {
		buf := src[i]
		if buf == '{' || buf == '[' || buf == '(' {
			push(&node{
				Elem:  buf,
				Index: i,
			})
			continue
		}

		if buf == ')' || buf == ']' || buf == '}' {
			pre := peek()
			if pre == nil {
				//空栈遇到右括号 将其进栈
				push(&node{
					Elem:  buf,
					Index: i,
				})
				continue
			}
			if buf == ')' && pre.Elem == '(' {
				pop()
				setIndexArrTrue(pre.Index, i+1)
				continue
			}
			if buf == ']' && pre.Elem == '[' {
				pop()
				setIndexArrTrue(pre.Index, i+1)
				continue
			}
			if buf == '}' && pre.Elem == '{' {
				pop()
				setIndexArrTrue(pre.Index, i+1)
				continue
			}
			//单独一个右括号 将其进栈
			push(&node{
				Elem:  buf,
				Index: i,
			})
		}
		//不是左右括号 直接跳过

	}

	//现在已经将匹配括号内的元素对应的 在indexArr中全部置true 寻找最长的即可
	// var from, to int = -1, -1
	// var buf1, buf2 int = -1, -1
	maxLen := 0
	bufLen := 0
	i := 0
	for {
		//遍历结束 比较一次
		if i >= len(src) {
			if bufLen > maxLen {
				maxLen = bufLen
			}
			break
		}

		if indexArr[i] == true {
			bufLen++
			i++
			continue
		} else {
			if bufLen > maxLen {
				maxLen = bufLen
			}
			bufLen = 0
			i++
		}

	}

	return maxLen

}

//将下下标数组 [from, to) 置true
func setIndexArrTrue(from, to int) {
	for i := from; i < to; i++ {
		indexArr[i] = true
	}
}

func initStack(cap int) {
	stack = make([]*node, cap)
	top = -1
}

func push(elem *node) {
	top++
	stack[top] = elem
}

func pop() bool {
	if top <= -1 {
		return false
	}
	top--
	return true
}

func peek() *node {
	if top == -1 {
		return nil
	}
	return stack[top]
}
