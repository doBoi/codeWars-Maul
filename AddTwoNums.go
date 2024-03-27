package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// Done!!
//LEET CODE

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs1, rs2, rv1, rv2 string
	//get value
	for i := l1; i != nil; i = i.Next {
		rs1 += strconv.Itoa(i.Val)
	}
	for i := l2; i != nil; i = i.Next {
		rs2 += strconv.Itoa(i.Val)

	}
	// reverse value
	for i := len(rs1) - 1; i >= 0; i-- {
		rv1 += string(rs1[i])
	}
	for i := len(rs2) - 1; i >= 0; i-- {
		rv2 += string(rs2[i])
	}
	r1, _ := new(big.Int).SetString(rv1, 10)
	r2, _ := new(big.Int).SetString(rv2, 10)
	rs := new(big.Int).Add(r1, r2).Text(10)
	r, _ := strconv.Atoi(string(rs[0]))
	h1 := &ListNode{Val: r}
	for i := 1; i <= len(rs)-1; i++ {
		r, _ = strconv.Atoi(string(rs[i]))
		h1 = &ListNode{Val: int(r), Next: h1}
	}
	return h1
}

func main() {
	n1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	n2 := []int{5, 6, 4}
	// var l1, l2 []ListNode
	l1 := &ListNode{Val: n1[len(n1)-1]}
	l2 := &ListNode{Val: n2[len(n2)-1]}

	for i := len(n1) - 2; i >= 0; i-- {
		l1 = &ListNode{Val: n1[i], Next: l1}
	}

	for i := len(n2) - 2; i >= 0; i-- {
		l2 = &ListNode{Val: n2[i], Next: l2}
	}

	fmt.Println(addTwoNumbers(l1, l2))
	// fmt.Println("l1", l1.Next)

	// for i := l1; i.Next != nil; i = i.Next {
	// 	fmt.Println(i)
	// }

}

func backup(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs1, rs2, rv1, rv2 string
	//get value
	for i := l1; i != nil; i = i.Next {
		rs1 += strconv.Itoa(i.Val)
	}
	for i := l2; i != nil; i = i.Next {
		rs2 += strconv.Itoa(i.Val)

	}
	// reverse value
	for i := len(rs1) - 1; i >= 0; i-- {
		rv1 += string(rs1[i])
	}
	for i := len(rs2) - 1; i >= 0; i-- {
		rv2 += string(rs2[i])
	}
	r1, _ := strconv.ParseUint(rv1, 10, 64)
	r2, _ := strconv.ParseUint(rv2, 10, 64)
	rs := strconv.FormatUint(r1+r2, 10)
	r, _ := strconv.Atoi(string(rs[0]))
	h1 := &ListNode{Val: r}
	for i := 1; i <= len(rs)-1; i++ {
		r, _ = strconv.Atoi(string(rs[i]))
		h1 = &ListNode{Val: int(r), Next: h1}
	}
	return h1
}
