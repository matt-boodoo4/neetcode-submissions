/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // key points 
 // looping through linked list with a pointer ( curr )
 // changing pointers - keep a tmp of next 
 // reversing means next points to prev
 // [0 -> 1 -> 2]
 // [0 -> nil & 1 -> 2]
 // [ 1 -> 0 -> nil]

func reverseList(head *ListNode) *ListNode {
	curr := head 
	var prev *ListNode
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev 
		prev = curr 
		curr = tmp
	
	}
	return prev
}
