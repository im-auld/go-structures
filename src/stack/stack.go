package main

import (
	"bytes"
	"fmt"
)

// Node for use in a Stack
type Node struct {
	value int
	next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("<Node: %v>", n.value)
}

// Stack - A Last In, First Out (LIFO) structure
type Stack struct {
	top *Node
}

func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for current := s.top; current != nil; current = current.next {
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

// Push a ``Node`` on to the ``Stack``
func (s *Stack) Push(node *Node) {
	node.next = s.top
	s.top = node
}

// Pop the top off the ``Stack``
func (s *Stack) Pop() *Node {
	popped := s.top
	s.top = s.top.next
	return popped
}

func main() {
	node1 := &Node{10, nil}
	node2 := &Node{15, nil}
	node3 := &Node{20, nil}
	stack := &Stack{node1}
	stack.Push(node2)
	stack.Push(node3)
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack)

}
