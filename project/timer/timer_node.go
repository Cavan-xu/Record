package timingwhell

import (
	"awesomeProject/project/common"
	"awesomeProject/project/snowflake"
)

type TimerNode struct {
	uniqueId       int64
	totalCount     int32
	currentCount   int32
	interval       int32
	nextTickTime   int64
	milliSecondPos int32
	prev           *TimerNode
	next           *TimerNode
	lst            *TimerList
}

func NewTimerNode(totalCount, interval int32) *TimerNode {
	return &TimerNode{
		uniqueId:       snowflake.New(1).GenerateId().Int64(),
		totalCount:     totalCount,
		currentCount:   0,
		nextTickTime:   common.GetCurMillionSeconds() + int64(interval),
		milliSecondPos: 0,
		interval:       interval,
	}
}

func (node *TimerNode) GetUniqueId() int64 {
	return node.uniqueId
}

func (node *TimerNode) GetLst() *TimerList {
	return node.lst
}

func (node *TimerNode) SetPrev(n *TimerNode) {
	node.prev = n
}

func (node *TimerNode) SetNext(n *TimerNode) {
	node.next = n
}

func (node *TimerNode) SetLst(lst *TimerList) {
	node.lst = lst
}

func (node *TimerNode) SetMilliSecondPos(milliSecondPos int32) {
	node.milliSecondPos = milliSecondPos
}

func (node *TimerNode) GetNext() *TimerNode {
	return node.next
}

func (node *TimerNode) GetPrev() *TimerNode {
	return node.prev
}

func (node *TimerNode) GetInterval() int32 {
	return node.interval
}
