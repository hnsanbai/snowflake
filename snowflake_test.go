package snowflake 

import (
	"sync"
	"testing"
	"runtime"
)

func TestGetSnowFlakeID(t *testing.T) {
	runtime.GOMAXPROCS(4)
	
	var wg sync.WaitGroup
	var m sync.Map
	wg.Add(4)
	for gs := 0; gs < 4; gs++ {
		go func() {
			for i := 0; i < 60; i++ {
				id := GetSnowFlakeID(10, 10)
				_, load := m.LoadOrStore(id, i)
				if load {
					t.Fatalf("Repeated ID: ", id)
				}
			}
			wg.Done()
		}()
	}
	
	wg.Wait()
	var i int64 = 0
	m.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	
	if i != 240 {
		t.Errorf("expected map length: %d", i)
	}
}