/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    list_aux := head
    nodeAddresses := make(map[ListNode]*ListNode)

    for list_aux != nil {
        if _, exists := nodeAddresses[*list_aux]; exists {
            return true
        }
        nodeAddresses[*list_aux] = list_aux
        list_aux = list_aux.Next
    }
    return false
}
