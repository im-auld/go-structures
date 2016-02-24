package main

import (
	"bytes"
	"fmt"
)

// Queuable is an interface to define a queuable object
type Queuable interface {
	Next()
}

// Node a node that can be queued
type Node struct {
	value interface{}
	next  *Node
}

// Next gets the next object
func (n *Node) Next() *Node {
	return n.next
}

// Job - A job for the queue
type Job struct {
	instruction string
	next        *Job
}

// Next gets the next object
func (j *Job) Next() *Job {
	return j.next
}

// Queue ...
type Queue struct {
	head *Node
	size int
}

func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for current := q.head; current != nil; current = current.Next() {
		if current.next == nil {
			buffer.WriteString(fmt.Sprintf("%v", current.value))
			break
		} else {
			buffer.WriteString(fmt.Sprintf("%v, ", current.value))
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Init ...
func (q *Queue) Init() *Queue {
	q.size = 0
	return q
}

// NewQueue ...
func NewQueue() *Queue {
	return new(Queue).Init()
}

// Enqueue - Add a Queueable to the Queue
func (q *Queue) Enqueue(node *Node) {
	if q.head == nil {
		q.head = node
	} else {
		for current := q.head; ; current = current.Next() {
			if current.Next() == nil {
				current.next = node
				break
			}
		}
	}
}

// Dequeue - Remove a Queueable form the Queue
func (q *Queue) Dequeue() *Node {
	result := q.head
	q.head = q.head.Next()
	return result
}

// Size - Return the size of the Queue
func (q *Queue) Size() int {
	return q.size
}

func main() {
	queue := NewQueue()
	node1 := &Node{1, nil}
	node2 := &Node{2, nil}
	node3 := &Node{3, nil}
	queue.Enqueue(node1)
	queue.Enqueue(node2)
	queue.Enqueue(node3)
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
}
