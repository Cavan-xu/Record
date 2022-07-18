package list

import "sync"

/*
	双向链表
*/

type node struct {
	val  int32
	prev *node
	next *node
	list *list
}

type list struct {
	sync.Mutex
	tail *node
}

func NewList() *list {
	tail := &node{}
	tail.prev = tail
	tail.next = tail

	return &list{tail: tail}
}

func (l *list) GetTail() *node {
	l.Lock()
	defer l.Unlock()

	return l.tail
}

func (l *list) GetFirst() *node {
	l.Lock()
	defer l.Unlock()

	return l.tail.next
}

func (l *list) GetLast() *node {
	l.Lock()
	defer l.Unlock()

	return l.tail.prev
}

func (l *list) GetPrev(node *node) *node {
	l.Lock()
	defer l.Unlock()

	if node.list != l {
		return l.tail
	}

	return node.prev
}

func (l *list) GetNext(node *node) *node {
	l.Lock()
	defer l.Unlock()

	if node.list != l {
		return l.tail
	}

	return node.next
}

func (l *list) PushFront(node *node) {
	l.Lock()
	defer l.Unlock()

	node.list = l
	node.prev = l.tail
	node.next = l.tail.next
	l.tail.next.prev = node
	l.tail.next = node
}

func (l *list) PushBack(node *node) {
	l.Lock()
	defer l.Unlock()

	node.list = l
	node.prev = l.tail.prev
	node.next = l.tail
	l.tail.prev.next = node
	l.tail.prev = node
}

func (l *list) Remove(node *node) {
	l.Lock()
	defer l.Unlock()

	if node.list != l {
		return
	}

	node.list = nil
	node.prev.next = node.next
	node.next.prev = node.prev
}
