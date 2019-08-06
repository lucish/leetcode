package two

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root1 := l1
	root2 := l2

	var isOverFolw bool //进位

	for {
		//检查上一位的进位
		if isOverFolw {
			l1.Val += 1
			isOverFolw = false
		}

		l1.Val += l2.Val

		//检查本位进位
		if l1.Val >= 10 {
			l1.Val -= 10
			isOverFolw = true
		}

		l2.Val = l1.Val

		if l1.Next == nil || l2.Next == nil {
			//两个一样长 新建一个进位
			if l1.Next == nil && l2.Next == nil {
				if isOverFolw {
					l1.Next = &ListNode{
						Val:  1,
						Next: nil,
					}
				}
				return root1
			}

			//一个链表较长

			//l2比较长 向l2进位并判断
			if l1.Next == nil {
				if isOverFolw {
					for {
						//进位并判断下一位是否溢出
						l2 = l2.Next
						l2.Val += 1
						if l2.Val < 10 {
							//未溢出可以返回
							break
						}
						//溢出归零并向上进位
						l2.Val = 0
						//上一位为空则新建并赋值1
						if l2.Next == nil {
							l2.Next = &ListNode{
								Val:  1,
								Next: nil,
							}
							break
						}
					}
				}
				return root2
			}

			if l2.Next == nil {
				//同上
				if isOverFolw {
					for {
						l1 = l1.Next
						l1.Val += 1
						if l1.Val < 10 {
							break
						}
						l1.Val = 0
						if l1.Next == nil {
							l1.Next = &ListNode{
								Val:  1,
								Next: nil,
							}
							break
						}
					}
				}
				return root1
			}

		}
		l1 = l1.Next
		l2 = l2.Next

	}

}
