/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 * Space: O(1)
 * Time: O(n)
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := false
	added := l1
	for {
		if extra {
			l1.Val++
			extra = false
		}

		l1.Val += l2.Val
		if l1.Val >= 10 {
			extra = true
			l1.Val -= 10
		}

		if l1.Next == nil && l2.Next == nil {
			if extra {
				l1.Next = &ListNode{1, nil}
			}
			return added
		}

		if l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			l2.Next = nil
		}

		l1 = l1.Next
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}
	}
}

// @lc code=end

