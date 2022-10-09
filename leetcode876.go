package main


type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	for i := 0; i < length/2; i++ {
		head = head.Next
	}
	return head
}

func main() {

}