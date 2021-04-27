package uuid

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
snowflake = 1bit（符号位） + 41bit（时间戳）+ 10bit（工作机器）+ 12bit（序列号）
1bit（符号位）：符号位，不使用（0为正数，1为负数）
41bit（时间戳）：2^41-1个数字代表69年，所以设置发号起始时间最好为发号器首次运行时间
10bit（工作机器）：也可以分为5bit（数据中心ID）+ 5bit（工作机器ID）或其它规则
12bit（序列号）：2^12-1个数字，总共4095，同一毫秒同一机器节点可以并发产生4095个不同Id
*/

// global var
var sequence int = 0
var lastTime int = -1

// every segment bit
var workerIdBits = 5
var datacenterIdBits = 5
var sequenceBits = 12

// every segment max number
var maxWorkerId int = -1 ^ (-1 << workerIdBits)
var maxDatacenterId int = -1 ^ (-1 << datacenterIdBits)
var maxSequence int = -1 ^ (-1 << sequenceBits)

// bit operation shift
var workerIdShift = sequenceBits
var datacenterShift = workerIdBits + sequenceBits
var timestampShift = datacenterIdBits + workerIdBits + sequenceBits

type Snowflake struct {
	datacenterId int
	workerId     int
	epoch        int
	mt           *sync.Mutex
}

func NewSnowflake(datacenterId int, workerId int, epoch int) (*Snowflake, error) {
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		return nil, errors.New(fmt.Sprintf("datacenterId cant be greater than %d or less than 0", maxDatacenterId))
	}
	if workerId > maxWorkerId || workerId < 0 {
		return nil, errors.New(fmt.Sprintf("workerId cant be greater than %d or less than 0", maxWorkerId))
	}
	if epoch > getCurrentTime() {
		return nil, errors.New(fmt.Sprintf("epoch time cant be after now"))
	}
	sf := Snowflake{datacenterId, workerId, epoch, new(sync.Mutex)}
	return &sf, nil
}

func (sf *Snowflake) getUniqueId() int {
	sf.mt.Lock()
	defer sf.mt.Unlock()
	// get current time
	currentTime := getCurrentTime()
	// compute sequence
	if currentTime < lastTime { // occur clock back
		// panic or wait,wait is not the best way.can be optimized.
		currentTime = waitUntilNextTime(lastTime)
		sequence = 0
	} else if currentTime == lastTime { // at the same time(micro-second)
		sequence = (sequence + 1) & maxSequence
		if sequence == 0 { // overflow max num,wait next time
			currentTime = waitUntilNextTime(lastTime)
		}
	} else if currentTime > lastTime { // next time
		sequence = 0
		lastTime = currentTime
	}
	// generate id
	return (currentTime-sf.epoch)<<timestampShift | sf.datacenterId<<datacenterShift |
		sf.workerId<<workerIdShift | sequence
}

func waitUntilNextTime(last int) int {
	currentTime := getCurrentTime()
	for currentTime <= last {
		time.Sleep(1 * time.Millisecond) // sleep micro second
		currentTime = getCurrentTime()
	}
	return currentTime
}

func getCurrentTime() int {
	return int(time.Now().UnixNano() / 1e6) // micro second
}
