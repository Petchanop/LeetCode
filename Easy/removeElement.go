package main

import "fmt"

/**
 * Definition for singly-linked list.
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(value int) *ListNode {
	var n ListNode
	n.Val = value
	n.Next = nil
	return &n
}

func ListIndex(head *ListNode, List *ListNode) int {
	for i := 1; head != nil; i++ {
		if head.Val == List.Val {
			return (i)
		}
		head = head.Next
	}
	return (0)
}

func delElements(head *ListNode) {
	for head != nil {
		head = head.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	var new *ListNode = head
	var pre *ListNode = head
	var tmp *ListNode = nil
	j := 1
	pre = nil
	for head.Next != nil {
		if head.Val == val {
			if ListIndex(new, head) > 1 {
				tmp = head.Next
				del := head
				delElements(del)
				pre.Next = tmp
				head = tmp
				continue
			} else {
				new = new.Next
				pre = new
			}
		}
		pre = head
		head = head.Next
		j++
	}
	if head.Val == val {
		if ListIndex(new, head) != 1 {
			del := head
			delElements(del)
			pre.Next = nil
		} else {
			new = new.Next
		}
	}
	return (new)
}

func main() {
	var test *ListNode = NewNode(1)
	test.Next = NewNode(1)
	test.Next.Next = NewNode(1)
	test.Next.Next.Next = NewNode(1)
	test.Next.Next.Next.Next = NewNode(1)
	test.Next.Next.Next.Next.Next = NewNode(1)
	test.Next.Next.Next.Next.Next.Next = NewNode(1)
	test.Next.Next.Next.Next.Next.Next.Next = nil
	new := test
	// for test != nil {
	// 	fmt.Println(test.Val)
	// 	test = test.Next
	// }
	rem := removeElements(new, 1)
	for rem != nil {
		fmt.Println(rem.Val)
		rem = rem.Next
	}
}
