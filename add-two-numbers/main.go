// https://leetcode.com/problems/add-two-numbers/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var dl1 = []int{0}
var dl2 = []int{0}

var l1 = &ListNode{}
var l2 = &ListNode{}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1s := l1
	l2s := l2

	r := &ListNode{}
	rh := r
	for r != nil {
		if l1s != nil {
			r.Val += l1s.Val
			l1s = l1s.Next
		}
		if l2s != nil {
			r.Val += l2s.Val
			l2s = l2s.Next
		}
		if r.Val >= 10 {
			r.Next = &ListNode{Val: 1}
			r.Val -= 10
		} else if l1s != nil || l2s != nil {
			r.Next = &ListNode{Val: 0}
		}
		r = r.Next
	}

	return rh
}

func printL(l *ListNode) string {
	r := ""
	for h := l; h != nil; h = h.Next {
		r += fmt.Sprint(h.Val)
	}
	return r
}

func fill(l *ListNode, dt []int) {
	lh := l
	for i := 0; i < len(dt); i++ {
		lh.Val = dt[i]
		if i < len(dt)-1 {
			lh.Next = &ListNode{}
			lh = lh.Next
		}
	}

}
func main() {
	fill(l1, dl1)
	fill(l2, dl2)

	l3 := addTwoNumbers(l1, l2)
	fmt.Println(printL(l3))

}
