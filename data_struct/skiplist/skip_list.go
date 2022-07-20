package skiplist

import (
	"math"

	"awesomeProject/practical_skills/common"
)

/*
	跳跃表：通过创建多级索引，已空间换时间，解决链表访问 o(n) 时间复杂度的问题

	1
	1       5
	1   3   5   7   9
	1 2 3 4 5 6 7 8 9

	查找8：1 5 7 8
*/

const (
	PromoteRate = 0.5
)

type Node struct {
	val   int32
	left  *Node // 前
	right *Node // 后
	up    *Node // 上
	down  *Node // 下
}

func NewNode(val int32) *Node {
	return &Node{val: val}
}

type SkipList struct {
	tail        *Node   // 头节点尾节点合一
	promoteRate float64 // 新添加的节点晋升到索引层的概率
	maxLevel    int     // 索引层最大层，从零开始
}

func NewSkipList(rate float64) *SkipList {
	tail := NewNode(math.MaxInt32)
	tail.left = tail
	tail.right = tail

	return &SkipList{tail: tail, promoteRate: rate}
}

func (s *SkipList) Add(val int32) {
	preNode := s.findNode(val)
	if preNode.val == val {
		return
	}

	node := NewNode(val)
	s.appendAfter(preNode, node) // 插入到最底层的链表中

	curLevel := 0
	for s.promote() {
		if curLevel == s.maxLevel { // 当前层已为最高层，需要增加一层
			s.addLevel()
		}
		for preNode.up == nil { // 找到上一层的前置节点
			preNode = preNode.left
		}
		preNode = preNode.up
		upperNode := NewNode(val)
		s.appendAfter(preNode, upperNode)
		upperNode.down = node
		node.up = upperNode
		node = upperNode
		curLevel++
	}

}

func (s *SkipList) Search(val int32) *Node {
	node := s.findNode(val)
	if node.val == val {
		return node
	}

	return nil
}

func (s *SkipList) Delete(val int32) {

}

// 找到值对应的前置节点
func (s *SkipList) findNode(val int32) *Node {
	node := s.tail

	for {
		for node.right != s.tail && node.right.val <= val {
			node = node.right
		}
		if node.down == nil {
			break
		}
		node = node.down
	}

	return node
}

// 把节点插到前置节点后面
func (s *SkipList) appendAfter(preNode *Node, node *Node) {
	node.left = preNode
	node.right = preNode.right
	preNode.right.left = node
	preNode.right = node
}

func (s *SkipList) promote() bool {
	return common.RandomFloat() >= s.promoteRate
}

func (s *SkipList) addLevel() {
	s.maxLevel++

	tail := NewNode(math.MaxInt32)
	tail.left = tail
	tail.right = tail
	tail.down = s.tail
	s.tail.up = tail

	s.tail = tail
}

func (s *SkipList) removeLevel() {

}
