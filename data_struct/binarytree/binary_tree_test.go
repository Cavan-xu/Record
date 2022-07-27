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
	tree.InOrder(tree.root)

	res := tree.InOrderV2()
	t.Log(res)
}
