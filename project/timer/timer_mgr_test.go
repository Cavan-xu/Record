package timingwhell

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTimerMgr(t *testing.T) {
	mgr := NewTimerMgr()
	assert.NotNil(t, mgr)

	mgr.SetTimer(10, IntervalSecond, func() string {
		fmt.Println("triggerTimer")
		return ""
	})

	select {}
}
