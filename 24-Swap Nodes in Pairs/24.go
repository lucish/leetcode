package main

import (
	"fmt"
)

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

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 ??? 这是在练习数据结构吗?

func swapPairs(head *ListNode) *ListNode {

	pre := &ListNode{}
	next := &ListNode{}
	one := &ListNode{}
	two := &ListNode{}
	root := &ListNode{}
	//pre 为新链表头的空节点
	pre.Next = head
	root = pre
	one = head
	for {
		//剩余0个或1个 不在需要进行交换
		if one == nil {
			break
		}
		two = one.Next
		if two == nil {
			break
		}
		//剩余两个 需要交换 同时初始化next 可能为空
		next = two.Next
		//swap
		pre.Next = two
		two.Next = one
		one.Next = next
		//更新pre 与 one
		pre = one
		one = next
	}
	return root.Next
}

func main() {
	node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}

	buf := swapPairs(node1)

	for buf != nil {
		fmt.Println(buf.Val)
		buf = buf.Next
	}
}
