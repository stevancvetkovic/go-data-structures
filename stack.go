package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(elem int) {
	s.items = append(s.items, elem)
}

func (s *Stack) pop() int {
	stackLength := len(s.items)
	toRemove := s.items[stackLength-1]
	s.items = s.items[:stackLength-1]
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.push(3)
	myStack.push(5)
	myStack.push(7)
	myStack.push(9)
	fmt.Println(myStack)

	removed := myStack.pop()
	fmt.Printf("removed %d\n", removed)
	fmt.Println(myStack)

	removed = myStack.pop()
	fmt.Printf("removed %d\n", removed)
	fmt.Println(myStack)
}
