package short_url

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch         = 1700000000000 // 自定义纪元（毫秒）
	machineIDBits = 10
	seqBits       = 12
	maxSeq        = -1 ^ (-1 << seqBits)                     // 4095
	maxMachineID  = int64(-1) ^ (int64(-1) << machineIDBits) // 1023
)

type Snowflake struct {
	mu        sync.Mutex
	machineID int64
	lastStamp int64
	sequence  int64
}

func New(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, fmt.Errorf("machineID must be in [0, %d]", maxMachineID)
	}
	return &Snowflake{machineID: machineID}, nil
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - epoch

	if now == s.lastStamp {
		s.sequence = (s.sequence + 1) & maxSeq
		if s.sequence == 0 {
			for now <= s.lastStamp {
				now = time.Now().UnixMilli() - epoch
			}
		}
	} else {
		s.sequence = 0
	}
	s.lastStamp = now

	// 正确写法：明确每段的位置
	// |  41位时间戳  |  10位机器ID  |  12位序列号  |
	return (now << (machineIDBits + seqBits)) |
		(s.machineID << seqBits) |
		s.sequence
}
