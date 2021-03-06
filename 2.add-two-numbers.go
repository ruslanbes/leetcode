/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 * Time: O(max(m,n))
 * Space: O(1)
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

		if l2 != nil {
			l1.Val += l2.Val
		}

		if l1.Val >= 10 {
			extra = true
			l1.Val -= 10
		}

		if l1.Next == nil && (l2 == nil || l2.Next == nil) {
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
		if l2 != nil {
			l2 = l2.Next
		}
	}
}

// @lc code=end

