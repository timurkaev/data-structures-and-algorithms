package main

import "fmt"

type PriorityQueue struct {
	maxSize  int
	queArray []int64
	nItems   int
}

// Constructor
func NewPriorityQueue(s int) *PriorityQueue {
	return &PriorityQueue{
		maxSize:  s,
		queArray: make([]int64, s),
		nItems:   0,
	}
}

// Insert an element into the priority queue
func (pq *PriorityQueue) Insert(item int64) {
	var j int
	if pq.nItems == 0 { // If the queue is empty
		pq.queArray[pq.nItems] = item // Insert the item into cell 0
		pq.nItems++
	} else { // Else, the queue is not empty
		for j = pq.nItems - 1; j >= 0; j-- { // Reverse iteration
			if item > pq.queArray[j] { // If the new element is greater
				pq.queArray[j+1] = pq.queArray[j] // Move the current element up
			} else { // If the new element is less or equal
				break // Shift stops
			}
		}
		pq.queArray[j+1] = item // Insert the element
		pq.nItems++
	}
}

// Remove the minimum element from the priority queue
func (pq *PriorityQueue) Remove() int64 {
	pq.nItems--
	return pq.queArray[pq.nItems]
}

// Peek the minimum element in the priority queue
func (pq *PriorityQueue) PeekMin() int64 {
	return pq.queArray[pq.nItems-1]
}

// Check if the queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.nItems == 0
}

// Check if the queue is full
func (pq *PriorityQueue) IsFull() bool {
	return pq.nItems == pq.maxSize
}

func main() {
	thePQ := NewPriorityQueue(5)
	thePQ.Insert(30)
	thePQ.Insert(50)
	thePQ.Insert(10)
	thePQ.Insert(40)
	thePQ.Insert(20)

	for !thePQ.IsEmpty() {
		item := thePQ.Remove()
		fmt.Print(item, " ")
	}
	fmt.Println()
}
