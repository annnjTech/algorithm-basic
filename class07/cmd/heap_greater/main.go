package main

import (
	"algorithm-basic/class07/heap_greater"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {

	h := heap_greater.NewHeapGreater(func(a, b any) int {
		return a.(*Student).Age - b.(*Student).Age
	})
	s1 := &Student{1, "Alice", 3}
	s2 := &Student{2, "Bob", 2}
	s3 := &Student{3, "Charlie", 5}
	s4 := &Student{4, "David", 1}
	s5 := &Student{5, "Eve", 4}
	h.Push(s1)
	h.Push(s2)
	h.Push(s3)
	h.Push(s4)
	h.Push(s5)

	s4.Age = 6
	h.Resign(s4)
	//h.Push(3)

	//fmt.Println(h.IsEmpty())
	//fmt.Println(h.Size())
	//fmt.Println(h.Contains(3))
	//fmt.Println(h.Peek())
	for _, val := range h.GetAll() {
		fmt.Println(val)
	}
}
