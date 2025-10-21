package TwoPointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthLastNode(head *ListNode, n int) *ListNode {
	listAux := new(ListNode)
	listAux = head
	pont1 := 0
	for {
		if listAux.Next != nil {
			listAux = listAux.Next
			pont1 += 1
		} else {
			break
		}
	}

	pont2 := pont1 - n + 1
	if pont1 == n-1 {
		listAux = head.Next
		return listAux
	}
	listAux = head
	for i := 0; i < pont2-1; i++ {
		listAux = listAux.Next
	}
	listAux2 := listAux.Next
	listAux.Next = listAux2.Next
	listAux = head
	return listAux
}

// Solução com dois ponteiros
/*
func removeNthLastNode(head *ListNode, n int) *ListNode {

	right := head
	left := head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next

	return head
}

*/
