package main

import (
	"container/list"
	"fmt"
)

var (
	SIZE  = 5
	cache = make(map[int]node)
)

type node struct {
	val  int
	next *node
	prev *node
}

func main() {
	input := [...]int{}

}
