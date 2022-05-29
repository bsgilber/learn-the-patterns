package main

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

// make hash set
// move through link list
// if .Next takes us to a val we've seen, break an return true
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)

	for {
		if head == nil {
			return false
		}

		if m[head] > 0 {
			return true
		}

		m[head] = 1
		head = head.Next
	}
}

// tortoise and hare
// one pointer moves ahead of the other, if they meet up again then it's a cycle, otherwise it's not
func hasCycleTandH(head *ListNode) bool {
	if head == nil {
		return false
	}

	var tort *ListNode = head
	var hare *ListNode = head.Next

	for {
		if hare == nil {
			return false
		}

		if tort == hare {
			return true
		}

		tort = tort.Next

		if hare.Next == nil || hare.Next.Next == nil {
			return false
		}

		hare = hare.Next.Next
	}
}
