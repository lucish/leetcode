package twenty_one

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	//合并链表首部的空节点
	var res = &ListNode{}
	buf := res
	for {
		//如果l1已空
		if l1 == nil {
			//将剩下的l2链接到链表上
			for l2 != nil {
				buf.Next = l2
				l2 = l2.Next
				buf = buf.Next
			}
			break
		}
		//同上
		if l2 == nil {
			//将剩下的l1链接到链表上
			for l1 != nil {
				buf.Next = l1
				l1 = l1.Next
				buf = buf.Next
			}
			break
		}

		if l1.Val <= l2.Val {
			buf.Next = l1
			l1 = l1.Next
			buf = buf.Next
		} else {
			buf.Next = l2
			l2 = l2.Next
			buf = buf.Next
		}

	}
	return res.Next
}
