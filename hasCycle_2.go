/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    ahead := head

    for ahead != nil && ahead.Next != nil {
        ahead = ahead.Next.Next
        head = head.Next
        if head == ahead {
            return true
        }
    }

    return false
}
