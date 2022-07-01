package timingwhell

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTimerMgr(t *testing.T) {
	mgr := NewTimerMgr()
	assert.NotNil(t, mgr)
}

func TestTimerMgr_LessThanOneMinute(t *testing.T) {
	mgr := NewTimerMgr()
	mgr.SetTimer(10, IntervalSecond, func() string {
		fmt.Println("triggerTimer")
		return ""
	})

	select {}
}

func TestTimerMgr_MoreThanOneMinute(t *testing.T) {
	mgr := NewTimerMgr()
	mgr.SetTimer(10, IntervalMinute, func() string {
		fmt.Println("triggerTimer")
		return ""
	})

	select{}
}
