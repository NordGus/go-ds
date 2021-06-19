package ds

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	length int
	head *node 
	tail *node
}

func (l *LinkedList) Len() int {
	return l.length
}

func (l *LinkedList) Empty() bool {
	return l.empty()
}

func (l *LinkedList) AddHead(data int) bool {
	return l.addHead(data)
}

func (l *LinkedList) AddTail(data int) bool {
	return l.addTail(data)
}

func (l *LinkedList) RemoveHead() (int, error) {
	return l.removeHead()
}

func (l *LinkedList) RemoveTail() (int, error) {
	return l.removeTail()
}

func (l *LinkedList) PeekHead() (int, error) {
	return l.peekHead()
}

func (l *LinkedList) PeekTail() (int, error) {
	return l.peekTail()
}

func (l *LinkedList) ToString() string {
	return l.toString()
}

func (l *LinkedList) Display() {
	fmt.Println(l.ToString())
}

func (l *LinkedList) addHead(data int) bool {
	n := &node{data: data}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}

	l.length++
	return true
}

func (l *LinkedList) addTail(data int) bool {
	n := &node{data: data}

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}

	l.length++
	return true
}

func (l *LinkedList) removeHead() (int, error) {
	if (l.empty()) {
		return 0, errors.New("empty list")
	}

	v := l.head.data

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		r := l.head
		l.head = l.head.next
		r.next = nil
	}

	l.length--
	return v, nil
}

func (l *LinkedList) removeTail() (int, error) {
	if (l.empty()) {
		return 0, errors.New("empty list")
	}

	v := l.tail.data

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		p := l.getPrevious(l.tail)
		p.next = nil
		l.tail = p
	}

	l.length--
	return v, nil
}

func (l *LinkedList) peekHead() (int, error) {
	if (l.empty()) {
		return 0, errors.New("empty list")
	}

	return l.head.data, nil;
}

func (l *LinkedList) peekTail() (int, error) {
	if (l.empty()) {
		return 0, errors.New("empty list")
	}

	return l.tail.data, nil;
}

func (l *LinkedList) toString() string {
	output := ""
	n := l.head

	for n != nil {
		if n == l.tail {
			output += fmt.Sprintf("%v", n.data)
		} else {
			output += fmt.Sprintf("%v -> ", n.data)
		}
		n = n.next
	}

	return output
}

func (l *LinkedList) empty() bool {
	return l.length == 0
}

func (l *LinkedList) getPrevious(n *node) *node {
	p := l.head
	
	for p != nil {
		if (p.next == n) {
			return p
		}

		p = p.next
	}

	return nil
}