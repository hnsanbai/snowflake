package snowflake

import (
	"time"
	"sync"
	"sync/atomic"
)

const (
	maxint64	= 0x7fffffffffffffff
	maxint42	= 0x1ffffffffff
	max5bit		= 0x1f
	max12bit	= 0xfff
	starttime	= 1525397999234
)

var (
	lastSF int64
	lastTimeMil int64
	lock sync.Mutex
)

func GetSnowFlakeID(mid1, mid2 int64) int64 {
	lock.Lock()
	defer lock.Unlock()
	
	newTimeMil := getTimeMill()
	if newTimeMil < lastTimeMil {
		time.Sleep(time.Duration(lastTimeMil-newTimeMil) + time.Millisecond)
		newTimeMil = getTimeMill()
	}
	// use the same machine code on the same machine
	if newTimeMil == lastTimeMil && atomic.LoadInt64(&lastSF) > max12bit {
		time.Sleep(time.Duration(lastTimeMil-newTimeMil) + time.Millisecond)
	}
	if newTimeMil > lastTimeMil {
		atomic.StoreInt64(&lastSF, 0)
	}
	
	lastTimeMil = newTimeMil
	
	if (lastTimeMil-starttime) > maxint42 {
		panic("out of range")
	}
	sf1 := (lastTimeMil-starttime) << 22
	
	if mid1 > max5bit || mid1 < 0 || mid2 > max5bit || mid2 < 0 {
		panic("out of range")
	}
	sf21 := mid1 << 17
	sf22 := mid2 << 12
	sf3 := lastSF
	
	result := maxint64&sf1|sf21|sf22|sf3
	
	if result > maxint64 {
		panic("out of range")
	}
	
	atomic.AddInt64(&lastSF, 1)
	return result
}

func getTimeMill() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}