package main

import (
	"fmt"
	"time"
)

var (
	cache    = make(map[int]int)
	lencache = 9
)

// type node struct{
// 	val int
// 	next *node
// 	prev *node
// }

func main() {

	input := [...]int{1, 2, 4, 5, 7, 2, 9}
	fmt.Println("LRU cache impl")
	for i := 0; i < 5; i++ {
		set(i, i*i)
	}

	printmap(cache)
	for _, v := range input {
		fmt.Println("============", v, "============")
		fmt.Println(get(v))
		printmap(cache)
	}

}

func printmap(m map[int]int) {
	for k, v := range m {
		fmt.Println("KEY: ", k, "  VAL: ", v)
	}

}
func set(k, v int) {
	cache[k] = v

}
func get(n int) int {
	if !consists(cache, n) {
		update(cache, n)
		fmt.Println("ZZzZZZzZZZzZZZzZZZzzzZZZzzZZZzzZZZZZZZzzzzZZZZzzz...")
		time.Sleep(3000 * time.Millisecond)
		return -1
	} else {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>...")
		time.Sleep(1000 * time.Millisecond)
		return n
	}

}

func update(m map[int]int, n int) { //TODO error check

	delete(cache, lencache-1)
	for i := lencache - 2; i >= 0; i-- {
		cache[i+1] = cache[i]

	}
	cache[0] = n

}

func consists(m map[int]int, n int) bool {
	for _, val := range m {
		if val == n {
			return true
		}
	}
	return false

}
