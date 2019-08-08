package twenty_five

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	函数返回对子链表排序后的首节点
	节点数小于四个 无需排序 直接返回首节点
	节点数大于四个 则需要排序 对前四个元素进行排序 并将最后一个节点链接到 reverseKGroup(subList *ListNode, k int)   subList为除去首部四个节点后的子链表
*/

func reverseKGroup(head *ListNode, k int) *ListNode {

	buf := head

	var node []*ListNode = make([]*ListNode, k)
	for i := 0; i < k; i++ {
		//剩余节点数小于k个 返回 递归结束
		if buf == nil {
			return head
		}
		node[i] = buf
		buf = buf.Next
	}

	//翻转k个节点
	for i := k - 1; i > 0; i-- {
		node[i].Next = node[i-1]
	}
	//遍历后head 在第k个节点的下一个节点 即子链表的head  对其进行递归
	subListHead := reverseKGroup(buf, k)
	node[0].Next = subListHead
	return node[k-1]
}
