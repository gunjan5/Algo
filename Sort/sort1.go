//sort these differently using the sort package, also sort them in reverse order
package main

import (
	"fmt"
	"sort"
	"strings"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	if res:=strings.Compare(strings.ToLower(p[i]), strings.ToLower(p[j])); res==-1 {
		return true
	} else {
		return false
	}
}

func (p people) Swap(i, j int)  {
	p[i],p[j]=p[j],p[i]

	
}


func main() {

	studyGroup := people {"Zeno", "John", "Al", "jenny"}

	fmt.Println(studyGroup)

	sort.Sort(studyGroup)

	fmt.Println(studyGroup)
	

}