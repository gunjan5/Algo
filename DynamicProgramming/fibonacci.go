package main

import (
	"fmt"
	"time"
)

func main() {

	in := 49
	t0 := time.Now()
	out := Fibo(in)
	t1 := time.Now()

	fmt.Println("Fibo: ", out)
	fmt.Println("Time: ", t1.Sub(t0))

}

func Fibo(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibo(n-1) + Fibo(n-2)
	}

}
