package main

import (
	"fmt"
	"time"
)

var (
	cache = make(map[int]int)
)

func main() {
	n := 49
	t0 := time.Now()
	op := FastFibo(n)
	t1 := time.Now()
	fmt.Println("Result: ", op)
	fmt.Println("Time: ", t1.Sub(t0))

}

func FastFibo(in int) (out int) {

	if in == 0 {
		if cache[in] != 0 {
			return cache[in]
		}
		cache[in] = 0
		return 0
	} else if in == 1 {
		if cache[in] != 0 {
			return cache[in]
		}
		cache[in] = 1
		return 1
	} else {
		if cache[in] != 0 {
			return cache[in]
		}
		cache[in] = FastFibo(in-1) + FastFibo(in-2)
		return cache[in]
	}

}
