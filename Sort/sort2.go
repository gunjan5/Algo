//sort these differently using the sort package, also sort them in reverse order
package main

import (
	"fmt"
	"sort"
)



func main() {
	var s sort.StringSlice

	s=[]string{"Zeno", "John", "Al", "jenny"}

	fmt.Println(s)
	sort.Sort(s)
	fmt.Println(s)


}