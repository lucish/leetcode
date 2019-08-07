package ninteen

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var fast, slow *ListNode

	fast = head
	//快的先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	//如果fast已经走完 链表长度也为n  即删除第一个节点
	if fast == nil {
		return head.Next
	}
	//再走一步
	fast = fast.Next
	slow = head
	//一起走 直到快的走到最后一个node之后的nil 此时慢的在倒数第n个节点的前置节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	//将slow节点的next指向其下一节点的next 即为删除slow的下一个节点

	slow.Next = slow.Next.Next
	return head
}
