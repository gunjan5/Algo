package main

import (
	"fmt"
	_"os"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head,tail *Node //	head, tail *Node
}

func main() {

	l := List{}
	// n1 := Node{val: 4}
	// l.head = &n1
	// n2 := Node{val: 6}
	// n1.next = &n2
	// n3 := Node{val: 8}
	// n2.next = &n3
	// n4 := Node{val: 10}
	// n3.next = &n4
	// n4.next = nil

	//fmt.Println(n1.val, n1.next.val, n1.next.next.val, n1.next.next.next.val)
	_ = "breakpoint"
	l.addTail(5)
	l.addTail(15)
	l.addTail(25)

	fmt.Println(l.findLen())
	l.printList()
	//fmt.Println(l)

}

func (l List) findLen() (n int) {
	for l.head != nil {
		n++
		l.head=l.head.next
	}
	return
}

func (l *List) addTail(i int) (err error) {
	n := Node{val: i}
	if(l.head==l.tail && l.head==nil){
		l.head = &n
	}else {
	l.tail.next= &n
	} 
	l.tail = &n
	l.tail.next =nil

	return nil	
}

func (l List) printList() {
	for l.head != nil {
		fmt.Println(l.head.val)
		l.head=l.head.next
	}
}