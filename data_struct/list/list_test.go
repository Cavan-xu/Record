package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	list := NewList()

	assert.NotNil(t, list)
}

func Test_list_GetFirst(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	list.PushFront(node1)
	list.PushFront(node2)

	assert.Equal(t, node2, list.GetFirst())
}

func Test_list_GetLast(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	list.PushFront(node1)
	list.PushFront(node2)

	assert.Equal(t, node1, list.GetLast())
}

func Test_list_GetPrev(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	list.PushFront(node1)
	list.PushFront(node2)

	assert.Equal(t, node2, list.GetPrev(node1))
}

func Test_list_GetNext(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	list.PushFront(node1)
	list.PushFront(node2)

	assert.Equal(t, node1, list.GetNext(node2))
}

func Test_list_PushFront(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	list.PushFront(node1)
	list.PushFront(node2)

	assert.Equal(t, node2, list.GetFirst())
}

func Test_list_PushBack(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	list.PushBack(node1)
	list.PushBack(node2)

	assert.Equal(t, node2, list.GetLast())
}

func Test_list_Remove(t *testing.T) {
	list := NewList()

	node1 := &node{val: 1}
	node2 := &node{val: 2}
	node3 := &node{val: 3}
	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)

	assert.Equal(t, node2, list.GetPrev(node3))
	list.Remove(node2)
	assert.Equal(t, node1, list.GetPrev(node3))
}
