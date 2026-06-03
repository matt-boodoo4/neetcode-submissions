/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}
	cycle := false 
	visited := make(map[*ListNode]bool)
	slow := head 
	fast := head.Next
	for slow != nil && fast != nil {
		if visited[fast] {
			return true
		}
		visited[slow] = true
		slow = slow.Next 
		if fast.Next != nil {
			fast = fast.Next.Next
		}
	}
	return cycle
}
