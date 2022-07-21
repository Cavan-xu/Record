package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList_Add(t *testing.T) {
	skipList := NewSkipList(PromoteRate)

	skipList.Add(1)
	skipList.Add(2)
	skipList.Add(3)
	skipList.Add(4)
	skipList.Add(5)
	skipList.Add(6)
	skipList.Add(7)

	skipList.LevelPrint()
}

func TestSkipList_Search(t *testing.T) {
	skipList := NewSkipList(PromoteRate)

	skipList.Add(1)
	skipList.Add(2)
	skipList.Add(3)
	skipList.Add(4)
	skipList.Add(5)
	skipList.Add(6)
	skipList.Add(7)

	node := skipList.Search(1)
	assert.Equal(t, int32(1), node.val)

	node = skipList.Search(5)
	assert.Equal(t, int32(5), node.val)

	node = skipList.Search(8)
	assert.Nil(t, node)

	skipList.LevelPrint()
}

func TestSkipList_Delete(t *testing.T) {
	skipList := NewSkipList(PromoteRate)

	skipList.Add(1)
	skipList.Add(2)
	skipList.Add(3)
	skipList.Add(4)
	skipList.Add(5)
	skipList.Add(6)
	skipList.Add(7)

	ok := skipList.Delete(1)
	assert.Equal(t, true, ok)

	ok = skipList.Delete(1)
	assert.Equal(t, false, ok)

	ok = skipList.Delete(3)
	assert.Equal(t, true, ok)

	ok = skipList.Delete(10)
	assert.Equal(t, false, ok)

	skipList.LevelPrint()
}
