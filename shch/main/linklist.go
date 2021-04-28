package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//链表中间节点
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {

}
