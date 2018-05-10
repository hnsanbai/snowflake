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
		go func(gi int) {
			for i := 0; i < 60; i++ {
				id := GetSnowFlakeID(10, 10)
				_, load := m.LoadOrStore(id, struct{}{})
				if load {
					t.Fatalf("Repeated ID: ", id)
				}
			}
			wg.Done()
		}(gs)
	}
	
	wg.Wait()
	var i int64 = 0
	m.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	t.Log(i)
	if i != 240 {
		t.Errorf("expected map length: %d", i)
	}
}