package chapter04

import (
	"sync"
	"testing"
	"time"
)

func TestMyLock(t *testing.T) {
	var (
		wg   sync.WaitGroup
		lock = NewMyLock()
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {

			defer func() {
				wg.Done()
			}()

			if lock.TryLock() {
				time.Sleep(time.Second)
				lock.Unlock()
				t.Logf("-------- %d get lock and unlock", num)
			} else {
				t.Logf("%d can not get lock", num)
			}
		}(i)
	}

	wg.Wait()
}

func BenchmarkMyLock(b *testing.B) {
	lock := NewMyLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if lock.TryLock() {
				time.Sleep(time.Second)
				lock.Unlock()
			}
		}
	})
}
