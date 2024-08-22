package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enqueue(elem int) {
	q.items = append(q.items, elem)
}

func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.enqueue(3)
	myQueue.enqueue(7)
	myQueue.enqueue(5)
	myQueue.enqueue(1)
	fmt.Println(myQueue)

	removed := myQueue.dequeue()
	fmt.Printf("removed %d\n", removed)
	fmt.Println(myQueue)

	removed = myQueue.dequeue()
	fmt.Printf("removed %d\n", removed)
	fmt.Println(myQueue)
}