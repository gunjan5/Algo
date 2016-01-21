//sort these differently using the sort package, also sort them in reverse order
package main

import (
	"fmt"
	"sort"
)

type num int


func main() {

	//var n sort.IntSlice

	n:=[]int{7,4,8,2,-9,12,-23,32,3}
	
	sort.Sort(sort.IntSlice(n))

	fmt.Println(n)


}