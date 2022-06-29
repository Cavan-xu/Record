package timingwhell

import "sync"

type TimerList struct {
	sync.Mutex
	tail *TimerNode
}

func NewTimerList() *TimerList {
	node := &TimerNode{}
	node.SetPrev(node)
	node.SetNext(node)
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

	return lst.tail.GetNext()
}

func (lst *TimerList) GetNext(node *TimerNode) *TimerNode {
	lst.Lock()
	defer lst.Unlock()

	if node.GetLst() != lst {
		return lst.tail
	}

	return node.GetNext()
}

func (lst *TimerList) GetPrev(node *TimerNode) *TimerNode {
	lst.Lock()
	defer lst.Unlock()

	if node.GetLst() != lst {
		return lst.tail
	}

	return node.GetPrev()
}

func (lst *TimerList) PushFront(node *TimerNode) {
	lst.Lock()
	defer lst.Unlock()

	node.SetLst(lst)
	node.SetPrev(lst.tail)
	node.SetNext(lst.tail.GetNext())
	lst.tail.GetNext().SetPrev(node)
	lst.tail.SetNext(node)
}

func (lst *TimerList) PushBack(node *TimerNode) {
	lst.Lock()
	defer lst.Unlock()

	node.SetLst(lst)
	node.SetNext(lst.tail)
	node.SetPrev(lst.tail.GetPrev())
	lst.tail.GetPrev().SetNext(node)
	lst.tail.SetPrev(node)
}

func (lst *TimerList) Remove(node *TimerNode) {
	lst.Lock()
	defer lst.Unlock()

	if node.GetLst() == lst {
		node.SetLst(nil)
		node.prev.SetNext(node.GetNext())
		node.next.SetPrev(node.GetPrev())
	}
}
