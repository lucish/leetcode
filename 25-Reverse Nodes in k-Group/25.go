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
	//首部新建一个空节点
	root := &ListNode{}
	root.Next = head
	var node [4]*ListNode

	for i := 0; i < 4; i++ {
		//剩余节点数小于四个 返回 递归结束
		if head == nil {
			return head
		}
		node[i] = head
		head = head.Next
	}

	subListHead := 

}
