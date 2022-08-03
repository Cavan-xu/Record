package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	inOder := []int{1, 5, 6, 8, 9, 11, 12, 13}
	preOrder := []int{8, 5, 1, 6, 9, 12, 11, 13}

	tree := NewTree(inOder, preOrder)
	assert.NotNil(t, tree)
}

func TestTree_InOrder(t *testing.T) {
	inOder := []int{1, 5, 6, 8, 9, 11, 12, 13}
	preOrder := []int{8, 5, 1, 6, 9, 12, 11, 13}

	tree := NewTree(inOder, preOrder)
	res1 := tree.InOrder(tree.root, []int{})
	res2 := tree.InOrderV2()
	assert.Equal(t, res1, res2)
}

func TestTree_PreOrder(t1 *testing.T) {
	inOder := []int{1, 5, 6, 8, 9, 11, 12, 13}
	preOrder := []int{8, 5, 1, 6, 9, 12, 11, 13}

	tree := NewTree(inOder, preOrder)
	res1 := tree.PreOrder(tree.root, []int{})
	res2 := tree.PreOrderV2()
	assert.Equal(t1, res1, res2)
}

func TestTree_AfterOrder(t1 *testing.T) {
	inOder := []int{1, 5, 6, 8, 9, 11, 12, 13}
	preOrder := []int{8, 5, 1, 6, 9, 12, 11, 13}

	tree := NewTree(inOder, preOrder)
	res1 := tree.AfterOrder(tree.root, []int{})
	res2 := tree.AfterOrderV2()
	assert.Equal(t1, res1, res2)
}
