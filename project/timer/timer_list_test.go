package timingwhell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimerList_GetTail(t *testing.T) {
	timerList := NewTimerList()
	tail := timerList.GetTail()
	assert.NotNil(t, tail)
}

func TestTimerList_GetFirstNode(t *testing.T) {
	timerList := NewTimerList()
	firstNode := timerList.GetFirstNode()

	assert.NotNil(t, firstNode)
	assert.Equal(t, int64(0), firstNode.uniqueId)

	timerList.PushFront(&TimerNode{uniqueId: 1})
	timerList.PushFront(&TimerNode{uniqueId: 2})
	firstNode = timerList.GetFirstNode()

	assert.NotNil(t, firstNode)
	assert.Equal(t, int64(2), firstNode.uniqueId)
}

func TestTimerList_GetNext(t *testing.T) {
	timerList := NewTimerList()

	node1 := &TimerNode{uniqueId: 1}
	node2 := &TimerNode{uniqueId: 2}
	node3 := &TimerNode{uniqueId: 3}

	timerList.PushFront(node1)
	timerList.PushFront(node2)
	timerList.PushFront(node3)

	next := timerList.GetNext(node2)
	assert.Equal(t, next, node1)
}

func TestTimerList_GetPrev(t *testing.T) {
	timerList := NewTimerList()

	node1 := &TimerNode{uniqueId: 1}
	node2 := &TimerNode{uniqueId: 2}
	node3 := &TimerNode{uniqueId: 3}

	timerList.PushFront(node1)
	timerList.PushFront(node2)
	timerList.PushFront(node3)

	prev := timerList.GetPrev(node2)
	assert.Equal(t, prev, node3)
}

func TestTimerList_PushFront(t *testing.T) {
	timerList := NewTimerList()

	node1 := &TimerNode{uniqueId: 1}
	node2 := &TimerNode{uniqueId: 2}
	node3 := &TimerNode{uniqueId: 3}

	timerList.PushFront(node1)
	timerList.PushFront(node2)
	timerList.PushFront(node3)

	tail := timerList.GetTail()
	assert.Equal(t, tail.next, node3)
}

func TestTimerList_PushBack(t *testing.T) {
	timerList := NewTimerList()

	node1 := &TimerNode{uniqueId: 1}
	node2 := &TimerNode{uniqueId: 2}
	node3 := &TimerNode{uniqueId: 3}

	timerList.PushBack(node1)
	timerList.PushBack(node2)
	timerList.PushBack(node3)

	tail := timerList.GetTail()
	assert.Equal(t, tail.prev, node3)
}

func TestTimerList_Remove(t *testing.T) {
	timerList := NewTimerList()

	node1 := &TimerNode{uniqueId: 1}
	node2 := &TimerNode{uniqueId: 2}
	node3 := &TimerNode{uniqueId: 3}

	timerList.PushFront(node1)
	timerList.PushFront(node2)
	timerList.PushFront(node3)

	next := timerList.GetNext(node3)
	assert.Equal(t, next, node2)

	timerList.Remove(node2)

	next = timerList.GetNext(node3)
	assert.Equal(t, next, node1)
}
