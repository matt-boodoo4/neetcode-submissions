/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // merge two sorted
 // sorted is the key - next is always greater than curr 
 // can do this iteratively.
 // use dummy list to do the merging ( new dummy head that's nil )
 // if l1.Val < l2.Val -> insert l1 before l2 
 // keep track of next pointers when merging 
 // can use a separate pointer (curr1 & curr2 ) to traverse 
 // if there's a remainder on either then we set them at the end. 

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    newHead := &ListNode{}
	currHead := newHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmp := list1.Next // keep the next reference
			currHead.Next = list1 // move list1 to next newpointer 
			list1 = tmp
		} else {
			tmp := list2.Next
			currHead.Next = list2
			list2  = tmp
		}
		currHead = currHead.Next
		
	}
	
	if list1 != nil {
		currHead.Next = list1
	}
	if list2 != nil {
		currHead.Next = list2
	}
	return newHead.Next
 }
