package timingwhell

import "sync"

type TimerList struct {
	sync.Mutex
	tail *TimerNode
}

func NewTimerList() *TimerList {
	node := &TimerNode{}
	node.prev = node
	node.next = node
	return &TimerList{tail: node}
}

func (lst *TimerList) GetTail() *TimerNode {
	lst.Lock()
	defer lst.Unlock()

	return lst.tail
}

func (lst *TimerList) GetFirstNode() *TimerNode {
	lst.Lock()
	defer lst.Unlock()

	return lst.tail.next
}

func (lst *TimerList) GetNext(node *TimerNode) *TimerNode {
	lst.Lock()
	defer lst.Unlock()

	if node.lst != lst {
		return lst.tail
	}

	return node.next
}

func (lst *TimerList) GetPrev(node *TimerNode) *TimerNode {
	lst.Lock()
	defer lst.Unlock()

	if node.lst != lst {
		return lst.tail
	}

	return node.prev
}

func (lst *TimerList) PushFront(node *TimerNode) {
	lst.Lock()
	defer lst.Unlock()

	node.lst = lst
	node.prev = lst.tail
	node.next = lst.tail.next
	lst.tail.next.prev = node
	lst.tail.next = node
}

func (lst *TimerList) PushBack(node *TimerNode) {
	lst.Lock()
	defer lst.Unlock()

	node.lst = lst
	node.next = lst.tail
	node.prev = lst.tail.prev
	lst.tail.prev.next = node
	lst.tail.prev = node
}

func (lst *TimerList) Remove(node *TimerNode) {
	lst.Lock()
	defer lst.Unlock()

	if node.lst != lst {
		return
	}

	node.lst = nil
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lst *TimerList) UnlockRemove(node *TimerNode) {
	if node.lst != lst {
		return
	}

	node.lst = nil
	node.prev.next = node.next
	node.next.prev = node.prev
}
