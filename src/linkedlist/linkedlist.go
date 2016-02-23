package main

import (
	"bytes"
	"fmt"
)

// Node ...
type Node struct {
	value interface{}
	next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("<Node: %v>", n.value)
}

// List ...
type List struct {
	head   *Node
	length int
}

// Init initializes an empty list
func (l *List) Init() *List {
	l.length = 0
	return l
}

// New returns an initialized list
func New() *List {
	return new(List).Init()
}

func (l *List) String() string {
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
func (l *List) Insert(node *Node) {
	if l.head != nil {
		node.next = l.head
		l.head = node
	} else {
		l.head = node
	}
	l.length++
}

// Search ..
func (l *List) Search(value int) *Node {
	for node := l.head; node != nil; node = node.next {
		if node.value == value {
			return node
		}
	}
	return nil
}

// Pop ...
func (l *List) Pop() *Node {
	node := l.head
	l.head = l.head.next
	node.next = nil
	l.length--
	return node
}

// Size ...
func (l *List) Size() int {
	return l.length
}

// Remove ...
func (l *List) Remove(node *Node) {
	var previous *Node
	for current := l.head; current != nil; current = current.next {
		if current == node {
			if current == l.head {
				l.head = current.next
			}
			if previous != nil {
				previous.next = current.next
			}
			current.next = nil
			l.length--
			break
		}
		previous = current
	}
}

func main() {
	node := Node{10, nil}
	node2 := Node{15, nil}
	node3 := Node{"Hello, world", nil}
	list := New()
	fmt.Println(list.Size())
	list.Insert(&node)
	list.Insert(&node2)
	list.Insert(&node3)
	fmt.Println(list.Size())
	fmt.Println(list)
	list.Remove(&node3)
	fmt.Println(list)
	fmt.Println(list.Size())
	// list.Insert(node2)
	// fmt.Println(list)
	// fmt.Println(fmt.Sprintf("Value 10 in List: %v", list.Search(10)))
	// fmt.Println(fmt.Sprintf("Value 15 in List: %v", list.Search(15)))
	// fmt.Println(fmt.Sprintf("Value 20 in List: %v", list.Search(20)))
	// fmt.Println(list.Size())
	// fmt.Println("list.pop() = ", list.Pop())
	// fmt.Println(list.Size())
	// list.Insert(&Node{20, nil})
	// fmt.Println(list)

}
