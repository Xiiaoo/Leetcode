package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	//节点数为0时
	if head == nil {
		return nil
	}

	//记录头节点地址
	cur := head

	//当前节点为空时跳出循环
	for cur.Next != nil {
		//如果当前节点Val值等于后继节点Val值，将当前节点指向后继节点下一个节点
		//如果当前节点Val值不等于后继节点Val值，将当前节点指向后继节点
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {

}
