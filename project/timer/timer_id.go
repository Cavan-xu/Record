package timingwhell

const (
	IntervalMillisecond int32 = 1 // 定时器毫秒为单位
	IntervalSecond            = 1000 * IntervalMillisecond
	IntervalMinute            = 60 * IntervalSecond
	IntervalHour              = 60 * IntervalMinute
	IntervalDay               = 24 * IntervalHour
	IntervalWeek              = 7 * IntervalDay
)
