/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    listAhead := head
    for listAhead != nil && listAhead.Next != nil {
        listAhead = listAhead.Next.Next
        head = head.Next
    }
    return head
}
