package main

import "fmt"

type Queue struct {
	maxSize  int     // The maximum queue size specified when the class was instantiated.
	queArray []int64 // An array used to store the elements of the queue.
	front    int     // The index of the first element in the queue.
	rear     int     // The index of the last element in the queue.
	nItems   int     // Number of elements in the queue
}
// Constructor
func NewQueue(s int) *Queue {
	return &Queue{
		maxSize:  s,
		queArray: make([]int64, s),
		front:    0,
		rear:     -1,
		nItems:   0,
	}
}

// Insert new element in Queue
func (q *Queue) Insert(j int64) {
	if q.rear == q.maxSize-1 {
		// Cyclic transfer
		q.rear = -1
	}
	q.rear++ // rear increment
	q.queArray[q.rear] = j
	q.nItems++ // Increasing the number of elements
}

// Removing an element from the queue
func (q *Queue) Remove() int64 {
	temp := q.queArray[q.front] // Inserting and increasing the front
	q.front++
	if q.front == q.maxSize {
		// Cyclic transfer
		q.front = 0
	}
	q.nItems-- // Reducing the number of elements
	return temp
}

// Reading an item at the beginning of the queue
func (q *Queue) PeekFront() int64 {
	return q.queArray[q.front]
}

// true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.nItems == 0
}

// true if the queue is full
func (q *Queue) IsFull() bool {
	return q.nItems == q.maxSize-1
}

// Number of items in the queue
func (q *Queue) Size() int {
	return q.nItems
}

func main() {
	theQueue := NewQueue(5) // queue of 5 cells

	// Inserting 4 elements
	theQueue.Insert(10)
	theQueue.Insert(20)
	theQueue.Insert(30)
	theQueue.Insert(40)

	// Extracting 3 elements
	theQueue.Remove()
	theQueue.Remove()
	theQueue.Remove()

	// Inserting 4 more elements with cyclic transfer
	theQueue.Insert(50)
	theQueue.Insert(60)
	theQueue.Insert(70)
	theQueue.Insert(80)

	// Extraction and output of all elements
	for !theQueue.IsEmpty() {
		n := theQueue.Remove()
		fmt.Print(n, " ")
	}
	fmt.Println()
}
