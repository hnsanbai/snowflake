package snowflake

import (
	"time"
	"sync"
)

const (
	maxint64	= 0x7fffffffffffffff
	max5bit		= 0x1f
	max12bit	= 0xfff
)

var (
	lastSF int64
	lastTimeMil int64
	lock sync.Mutex
)

func GetSnowFlakeID(mid1, mid2 int64) int64 {
	lock.Lock()
	defer lock.Unlock()
	
	if lastSF >= max12bit {
		lastSF = 0
	}
	
	newTimeMil := getTimeMill()
	if newTimeMil <= lastTimeMil {
		time.Sleep(time.Duration(lastTimeMil-newTimeMil) + time.Millisecond)
		newTimeMil = getTimeMill()
	}
	lastTimeMil = newTimeMil
	sf1 := lastTimeMil << 22
	
	if mid1 > max5bit || mid1 < 0 || mid2 > max5bit || mid2 < 0 {
		panic("out of range")
	}
	sf21 := mid1 << 17
	sf22 := mid2 << 12
	
	lastSF += 1
	sf3 := lastSF
	
	result := maxint64&sf1|sf21|sf22|sf3
	
	return result
}

func getTimeMill() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}