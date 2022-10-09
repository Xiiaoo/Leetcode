package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	if cur == nil {
		return cur
	}
	if cur.Val == val && cur.Next == nil {
		return nil
	} else {
		for cur.Next != nil {
			if cur.Val == val {
				cur.Val = cur.Next.Val
				cur.Next = cur.Next.Next
			} else if cur.Next.Val == val {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
		if cur.Val == val {
			head = nil
		}
	}
	return head
}

func main() {

}
