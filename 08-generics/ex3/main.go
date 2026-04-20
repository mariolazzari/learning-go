package main

import "fmt"

// Node holds a single element
type Node[T comparable] struct {
	value T
	next  *Node[T]
}

// LinkedList defines the behavior of a generic list
type LinkedList[T comparable] interface {
	Add(value T)
	Insert(value T, index int)
	Index(value T) int
}

// concrete implementation
type linkedList[T comparable] struct {
	head *Node[T]
	size int
}

// constructor (consigliato con interfacce)
func NewLinkedList[T comparable]() LinkedList[T] {
	return &linkedList[T]{}
}

// Add appends element at end
func (l *linkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}

	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	l.size++
}

// Insert adds element at index
func (l *linkedList[T]) Insert(value T, index int) {
	newNode := &Node[T]{value: value}

	if index <= 0 || l.head == nil {
		newNode.next = l.head
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	i := 0

	for curr.next != nil && i < index-1 {
		curr = curr.next
		i++
	}

	newNode.next = curr.next
	curr.next = newNode
	l.size++
}

// Index returns position of value or -1
func (l *linkedList[T]) Index(value T) int {
	curr := l.head
	i := 0

	for curr != nil {
		if curr.value == value {
			return i
		}
		curr = curr.next
		i++
	}

	return -1
}

func main() {
	var list LinkedList[int] = NewLinkedList[int]()

	list.Add(10)
	list.Add(20)
	list.Add(30)

	list.Insert(15, 1)

	fmt.Println(list.Index(20)) // 2
	fmt.Println(list.Index(99)) // -1
}
