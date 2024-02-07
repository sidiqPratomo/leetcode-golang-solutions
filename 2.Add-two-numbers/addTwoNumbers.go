package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	inc := false
	cur := res
	for l1 != nil || l2 != nil || inc {
		tmp := 0
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		if inc {
			tmp++
		}
		if tmp >= 10 {
			inc = true
			tmp -= 10
		} else {
			inc = false
		}
		cur.Next = &ListNode{Val: tmp}
		cur = cur.Next
	}
	return res.Next
}