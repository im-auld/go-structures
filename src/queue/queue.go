package main

import (
	"bytes"
	"fmt"
)

// Queuable is an interface to define a queuable object
type Queuable interface {
	Next() Queuable
	SetNext(Queuable)
}

// Base - A Base struct
type Base struct {
	next Queuable
}

// Next - Get the next object in the Queue
func (b *Base) Next() Queuable { return b.next }

// SetNext - Sets the next object
func (b *Base) SetNext(q Queuable) { b.next = q }

// Node - A plain value for the Queue
type Node struct {
	value interface{}
	*Base
}

func (n *Node) String() string {
	return fmt.Sprintf("<Node: %v>", n.value)
}

// Job - A job for the queue
type Job struct {
	instruction string
	*Base
}

func (j *Job) String() string {
	return fmt.Sprintf("<Job: %v>", j.instruction)
}

// Queue ...
type Queue struct {
	head Queuable
	size int
}

func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for current := q.head; current != nil; current = current.Next() {
		if current.Next() == nil {
			buffer.WriteString(fmt.Sprintf("%v", current))
			break
		} else {
			buffer.WriteString(fmt.Sprintf("%v, ", current))
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

// SetHead - Sets a new head for the Queue
func (q *Queue) SetHead(obj Queuable) {
	q.head = obj
}

// GetHead - Returns the head of the Queue
func (q *Queue) GetHead() Queuable {
	return q.head
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
func (q *Queue) Enqueue(node Queuable) {
	node.SetNext(q.head)
	q.head = node
	q.size++
}

// Dequeue - Remove a Queueable form the Queue
func (q *Queue) Dequeue() Queuable {
	result := q.GetHead()
	q.head = q.head.Next()
	q.size--
	return result
}

// Size - Return the size of the Queue
func (q *Queue) Size() int {
	return q.size
}

func main() {
	queue := NewQueue()
	node1 := &Node{1, &Base{}}
	job := &Job{"Learn more Go!", &Base{}}
	node3 := &Node{3, &Base{}}
	queue.Enqueue(node1)
	queue.Enqueue(job)
	queue.Enqueue(node3)
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
}
