package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head, tail *Node
}

func main() {
	n1 := Node(4)
	head = &n1
	n2 := Node(6)
	n1.next = n2
	tail = null
	n2.next = tail

}
