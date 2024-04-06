package linkedList

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	//head=list[i]
	//pre=list[i-1]
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		//pre=list[(i+2)-1]
		pre = head
		//head=list[(i+2)]
		head = next
	}
	return dummy.Next
}

// 递归版本
func swapPairsV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
