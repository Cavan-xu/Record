package timingwhell

import (
	"log"
	"runtime/debug"
	"sync"
	"time"

	"awesomeProject/project/common"
)

const (
	milliSecondPerMinute = 60 * 1000
	minutePerWeek        = 60 * 24 * 7
	milliSecondPerWeek   = milliSecondPerMinute * minutePerWeek
)

var (
	once     sync.Once
	timerMgr *TimerMgr
)

type TimerMgr struct {
	mu                sync.Mutex
	cursorMutex       sync.RWMutex
	milliSecondLists  []*TimerList
	minuteLists       []*TimerList
	timerMap          map[int64]*TimerNode
	milliSecondCursor int32
	minuteCursor      int32
	lastUpdateTime    int64
}

func NewTimerMgr() *TimerMgr {
	once.Do(func() {
		timerMgr = &TimerMgr{
			milliSecondLists:  make([]*TimerList, milliSecondPerMinute),
			minuteLists:       make([]*TimerList, minutePerWeek),
			timerMap:          make(map[int64]*TimerNode),
			milliSecondCursor: 0,
			minuteCursor:      0,
			lastUpdateTime:    common.GetCurMillionSeconds(),
		}
		for i := 0; i < milliSecondPerMinute; i++ {
			timerMgr.milliSecondLists[i] = NewTimerList()
		}
		for i := 0; i < minutePerWeek; i++ {
			timerMgr.minuteLists[i] = NewTimerList()
		}
	})

	go timerMgr.Work()
	return timerMgr
}

func (mgr *TimerMgr) SetTimer(totalCount, interval int32, delegate func() string) int64 {
	node := NewTimerNode(totalCount, interval, delegate)

	mgr.addTimerToMap(node)
	mgr.addTimerToList(node)

	return node.uniqueId
}

func (mgr *TimerMgr) RemoveTimer(uniqueId int64) {
	node, ok := mgr.getTimer(uniqueId)
	if !ok {
		return
	}

	mgr.removeFromMap(node)
	mgr.removeFromList(node)
}

func (mgr *TimerMgr) Work() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error: %v, stack: %s", err, string(debug.Stack()))
			go mgr.Work()
		}
	}()

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			mgr.update()
		}
	}
}

func (mgr *TimerMgr) addTimerToMap(node *TimerNode) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	mgr.timerMap[node.uniqueId] = node
}

func (mgr *TimerMgr) addTimerToList(node *TimerNode) {
	mgr.cursorMutex.RLock()
	defer mgr.cursorMutex.RUnlock()

	var lst *TimerList
	if node.interval < milliSecondPerMinute {
		pos := (mgr.milliSecondCursor + node.interval) % milliSecondPerMinute
		lst = mgr.milliSecondLists[pos]
	} else if node.interval <= milliSecondPerWeek {
		minutes := (mgr.milliSecondCursor + node.interval) / milliSecondPerMinute
		pos1 := (mgr.milliSecondCursor + node.interval) % milliSecondPerMinute
		pos2 := (mgr.minuteCursor + minutes) % minutePerWeek
		node.milliSecondPos = pos1
		lst = mgr.minuteLists[pos2]
	} else {
		mgr.removeFromMap(node)
		return
	}

	lst.PushBack(node)
}

func (mgr *TimerMgr) removeFromMap(node *TimerNode) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	delete(mgr.timerMap, node.uniqueId)
}

func (mgr *TimerMgr) removeFromList(node *TimerNode) {
	lst := node.lst
	if lst != nil {
		lst.Remove(node)
	}
}

func (mgr *TimerMgr) getTimer(uniqueId int64) (*TimerNode, bool) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	node, ok := mgr.timerMap[uniqueId]
	return node, ok
}

func (mgr *TimerMgr) update() {
	curMillisecond := common.GetCurMillionSeconds()
	crossMilliSecond := curMillisecond - mgr.lastUpdateTime
	if crossMilliSecond <= 0 {
		return
	}

	// 每次 tick 时间不能超过一分钟
	if crossMilliSecond >= milliSecondPerMinute {
		crossMilliSecond = milliSecondPerMinute - 1
	}

	mgr.lastUpdateTime += crossMilliSecond

	for i := 0; i < int(crossMilliSecond); i++ {
		mgr.cursorMutex.Lock()
		mgr.milliSecondCursor += 1
		if mgr.milliSecondCursor >= milliSecondPerMinute { // 一分钟过去了
			mgr.milliSecondCursor -= milliSecondPerMinute
			mgr.minuteCursor = (mgr.minuteCursor + 1) % minutePerWeek

			lst := mgr.minuteLists[mgr.minuteCursor]
			mgr.minuteToMilliSecondList(lst)
		}
		mgr.cursorMutex.Unlock()

		lst := mgr.milliSecondLists[mgr.milliSecondCursor]
		mgr.timeout(lst)
	}

}

func (mgr *TimerMgr) minuteToMilliSecondList(lst *TimerList) {
	lst.Lock()
	defer lst.Unlock()

	for node := lst.tail.next; node != lst.tail; {
		next := node.next
		lst.UnlockRemove(node)
		mgr.milliSecondLists[node.milliSecondPos].PushBack(node)
		node = next
	}
}

func (mgr *TimerMgr) timeout(lst *TimerList) {
	lst.Lock()
	defer lst.Unlock()

	for node := lst.tail.next; node != lst.tail; {
		next := node.next
		node.currentCount++
		mgr.triggerTimer(node)
		if node.currentCount < node.totalCount {
			lst.UnlockRemove(node)
			node.nextTickTime = mgr.lastUpdateTime + int64(node.interval)
			mgr.addTimerToList(node)
		} else {
			lst.UnlockRemove(node)
			mgr.removeFromMap(node)
		}
		node = next
	}
}

func (mgr *TimerMgr) triggerTimer(node *TimerNode) {
	node.onTimer()
}
