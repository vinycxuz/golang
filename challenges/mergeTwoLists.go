package challenges

// https://leetcode.com/problems/merge-two-sorted-lists/solutions/7149759/mergetwolists-by-vinycxuz-eabd/

/*
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	endList := new(ListNode)
	head := endList
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return endList.Next
}
*/
