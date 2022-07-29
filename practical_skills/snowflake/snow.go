package snowflake

// GenerateId 用于分布式场景下生成唯一 id
/*
  最高位是符号位，始终为 0，不可用。
  41 位的时间序列，精确到毫秒级，41位的长度可以使用69年。时间位还有一个很重要的作用是可以根据时间进行排序。
  10 位的机器标识，10位的长度最多支持部署1023个节点。
  12 位的计数序列号，序列号即一系列的自增id，可以支持同一节点同一毫秒生成多个ID序号，12位的计数序列号支持每个节点每毫秒产生4095个ID序号。

  符号位		41bit 时间戳									十位机器编号		 12位计数序号
  0 		00000000000000000000000000000000000000000	0000000000		000000000000
*/

import (
	"sync"
	"time"

	"awesomeProject/common"
)

const (
	workerBit uint8 = 10 // 机器位数
	stepBit   uint8 = 12 // 计数位数
	maxStep   int64 = -1 ^ (-1 << stepBit)

	timeShift = workerBit + stepBit // 时间左移位数
)

var (
	once sync.Once

	instance *SnowGenerator
)

type ID int64

func (id ID) Int64() int64 {
	return int64(id)
}

type SnowGenerator struct {
	mu    sync.Mutex
	epoch int64 // 单位毫秒

	node int64
	time int64
	step int64
}

// New SnowGenerator
func New(node int64) *SnowGenerator {
	once.Do(func() {
		startTime, _ := time.ParseInLocation(time.RFC3339, "2022-06-18T20:06:05+08:00", time.Local)
		epoch := common.GetTimeMillionSeconds(startTime)
		instance = &SnowGenerator{
			epoch: epoch,
			node:  node,
			time:  0,
			step:  0,
		}
	})
	return instance
}

func (sg *SnowGenerator) GenerateId() ID {
	sg.mu.Lock()
	defer sg.mu.Unlock()

	now := common.GetCurMillionSeconds()

	if sg.time == now { // 在同一个毫秒里面
		sg.step = (sg.step + 1) & maxStep
		if sg.step == 0 {
			for now <= sg.time {
				now = common.GetCurMillionSeconds()
			}
		}
	} else {
		sg.step = 0 // 重置计数序列号
	}

	sg.time = now

	return ID((now-sg.epoch)<<timeShift | sg.node<<workerBit | sg.step)
}
