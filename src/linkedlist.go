package main

import (
	"bytes"
	"fmt"
)

// Node ...
type Node struct {
	value int
	next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("<Node: %v>", n.value)
}

// List ...
type List struct {
	head *Node
}

func (l List) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for current := l.head; current != nil; current = current.next {
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

// Insert ...
func (l List) Insert(node *Node) bool {
	current := l.head
	for {
		if current.next == nil {
			current.next = node
			return true
		}
		current = current.next
	}
}

// Search ..
func (l List) Search(value int) bool {
	current := l.head
	for {
		if current.value == value {
			return true
		}
		if current.next == nil {
			return false
		}
		current = current.next
	}
}

// Pop ...
func (l List) Pop() *Node {
	if l.head.next != nil {
		node, l.head = l.head, l.head.next
	}
  else if l.head.next == nil {
    node = l.head
  }
  else {
    panic("Cannot pop form empty list")
  }
}

// Size ...
func (l List) Size() int {
	size := 0
	for node := l.head; node != nil; node = node.next {
		size++
	}
	return size
}

func main() {
	node := Node{10, nil}
	node2 := Node{15, nil}
	list := List{&node}
	// list2 := List{}

	list.Insert(&node2)
	fmt.Println(list)
	fmt.Println(fmt.Sprintf("Value 10 in List: %v", list.Search(10)))
	fmt.Println(fmt.Sprintf("Value 20 in List: %v", list.Search(20)))
	fmt.Println(list.Size())
}
