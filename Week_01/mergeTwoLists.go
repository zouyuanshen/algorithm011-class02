package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	result := newListNode

	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			newListNode.Next = l1.Next
			l1 = l1.Next

		} else {

			newListNode.Next = l2.Next
			l2 = l2.Next
		}
		newListNode = newListNode.Next
	}
	if l1 != nil {
		newListNode.Next = l1
	}

	if l2 != nil {
		newListNode.Next = l2
	}
	return result.Next
}
