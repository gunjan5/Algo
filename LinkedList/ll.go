package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node //	head, tail *Node
}

func main() {

	l := List{}
	n1 := Node{val: 4}
	l.head = &n1
	n2 := Node{val: 6}
	n1.next = &n2
	n3 := Node{val: 8}
	n2.next = &n3
	n4 := Node{val: 10}
	n3.next = &n4
	n4.next = nil

	fmt.Println(n1.val, n1.next.val, n1.next.next.val, n1.next.next.next.val)

}
