package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val, p1, p2 := 0, l1, l2
	for p1.Next != nil && p2.Next != nil {
		if p1.Val+p2.Val+val >= 10 {
			p1.Val = p1.Val + p2.Val + val - 10
			val = 1
		} else {
			p1.Val = p1.Val + p2.Val + val
			val = 0
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1.Next == nil && p2.Next == nil {
		if p1.Val+p2.Val+val >= 10 {
			p1.Val = p1.Val + p2.Val + val - 10
			newnode := new(ListNode)
			newnode.Val = 1
			p1.Next = newnode
		} else {
			p1.Val = p1.Val + p2.Val + val
		}
	} else if p1.Next == nil && p2.Next != nil {
		p1.Next = p2.Next
		if p1.Val+p2.Val+val >= 10 {
			p1.Val = p1.Val + p2.Val + val - 10
			val = 1
		} else {
			p1.Val = p1.Val + p2.Val + val
			val = 0
		}
		p2 = p2.Next
		for p2 != nil {
			if p2.Val+val == 10 {
				p2.Val = 0
				val = 1
			} else {
				p2.Val = p2.Val + val
				val = 0
			}
			if val == 1 && p2.Next == nil {
				lastnode := new(ListNode)
				lastnode.Val = 1
				p2.Next = lastnode
				break
			}
			p2 = p2.Next
		}
	} else if p1.Next != nil && p2.Next == nil {
		if p1.Val+p2.Val+val >= 10 {
			p1.Val = p1.Val + p2.Val + val - 10
			val = 1
		} else {
			p1.Val = p1.Val + p2.Val + val
			val = 0
		}
		p1 = p1.Next
		for p1 != nil {
			if p1.Val+val == 10 {
				p1.Val = 0
				val = 1
			} else {
				p1.Val = p1.Val + val
				val = 0
			}
			if val == 1 && p1.Next == nil {
				lastnode := new(ListNode)
				lastnode.Val = 1
				p1.Next = lastnode
				break
			}
			p1 = p1.Next
		}
	}
	return l1
}

func main() {

}
