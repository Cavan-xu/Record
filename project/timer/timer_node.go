package timingwhell

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
