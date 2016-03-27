package main

import (
	"fmt"
	"strings"
)

func Balance(s string) bool {
	opening := "[{("
	pair := map[string]string{
		"[": "]",
		"{": "}",
		"(": ")",
	}
	var stk Stack

	if len(s)%2 != 0 {
		return false
	} else {
		for _, v := range s {
			if strings.ContainsAny(opening, string(v)) {
				stk.Push(string(v))
			} else {
				if !(pair[stk.Pop()] == string(v)) {
					return false
				}
			}
		}
		if stk.isEmpty() {
			return true
		}
		return false
	}
}

type Stack struct {
	data []string
}

func (s Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(in string) {
	s.data = append(s.data, in)

}

func (s Stack) Print() string {
	k := ""
	for i, v := range s.data {
		k += fmt.Sprintf("\n Stack I:%d, V:%v", i, v)
	}
	return k

}
func (s Stack) Size() int {
	k := 0
	for range s.data {
		k++
	}
	return k
}

func (s *Stack) Pop() string {
	if s.isEmpty() {
		return ""
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// func main() {
// 	s := Stack{}

// 	s.Push("abc")
// 	s.Push("xyz")
// 	s.Push("qqq")
// 	s.Push("www")

// 	fmt.Println(s.isEmpty())
// 	fmt.Println(s.Print())

// 	s.Pop()
// 	s.Pop()
// 	fmt.Println(s.Print())

// 	fmt.Println(s.Size())
// }
