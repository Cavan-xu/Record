package snow_flake

/*
	算法描述：用于分布式场景下生成唯一 id
		最高位是符号位，始终为0，不可用。
		41位的时间序列，精确到毫秒级，41位的长度可以使用69年。时间位还有一个很重要的作用是可以根据时间进行排序。
		10位的机器标识，10位的长度最多支持部署1023个节点。
		12位的计数序列号，序列号即一系列的自增id，可以支持同一节点同一毫秒生成多个ID序号，12位的计数序列号支持每个节点每毫秒产生4095个ID序号。

		符号位		41bit 时间戳									十位机器编号		 12位计数序号
		  0 		00000000000000000000000000000000000000000	0000000000		000000000000
*/

import (
	"sync"
	"time"
)

var (
	startTime, _ = time.Parse(time.RFC3339, "2022-06-18T20:06:05+08:00")
	milliSecond  = int64(time.Millisecond)
	epoch        = startTime.UnixNano() / milliSecond

	workerBit uint8 = 10 // 机器位数
	stepBit   uint8 = 12 // 计数位数
	maxStep   int64 = -1 ^ (-1 << stepBit)

	timeShift = workerBit + stepBit // 	时间左移位数
)

type snowGenerator struct {
	sync.Mutex
	time int64
	node int64
	step int64
}

type ID int64

var (
	instance *snowGenerator
)

func init() {
	instance = &snowGenerator{
		time: 0,
		step: 0,
	}
}

func SetNode(node int64) {
	instance.node = node
}

func GenerateId() ID {
	instance.Lock()
	defer instance.Unlock()

	now := time.Now().UnixNano() / milliSecond

	if instance.time == now { // 在同一个毫秒里面
		instance.step = (instance.step + 1) & maxStep
		if instance.step == 0 {
			for now <= instance.time {
				now = time.Now().UnixNano() / milliSecond
			}
		}
	} else {
		instance.step = 0 // 重置计数序列号
	}

	instance.time = now

	return ID((now-epoch)<<timeShift | instance.node<<workerBit | instance.step<<stepBit)
}

func (id ID) Int64() int64 {
	return int64(id)
}
