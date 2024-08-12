package algoutil

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(vals ...int) *ListNode {
	dummyHead := &ListNode{}
	p := dummyHead
	for _, val := range vals {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		p.Next = node
		p = p.Next
	}
	ret := dummyHead.Next
	dummyHead.Next = nil
	return ret
}
func (h *ListNode) Print() {
	content := ""
	p := h
	for p.Next != nil {
		content += fmt.Sprintf("%d ", p.Next)
		p = p.Next
	}
	fmt.Println(content)
}
