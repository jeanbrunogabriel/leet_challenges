/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var reverseList *ListNode

    for head != nil {
        nextNode := head.Next
        head.Next = reverseList
        reverseList = head
        head = nextNode
    }
    return reverseList
}
