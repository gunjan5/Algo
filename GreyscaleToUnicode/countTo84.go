package main

import (
	"fmt"
	"reflect"
)

func main() {
	count := 84
	un := "\u25a4"

	fmt.Println("type is: ", reflect.TypeOf(un))
	fmt.Println(un)

	for i := 0; i <= count; i++ {
		fmt.Println(i, ":'',")

	}
}
