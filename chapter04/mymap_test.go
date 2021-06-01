package chapter04

import (
	"sync"
	"testing"
	"time"
)

func TestMyMapStore(t *testing.T) {
	m := NewMyMap()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {

		defer func() {
			wg.Done()
		}()

		for i := 1; i <= 10; i++ {
			m.Store(i, i)
		}

	}()

	go func() {

		defer func() {
			wg.Done()
		}()

		for i := 11; i <= 20; i++ {
			m.Store(i, i)
		}

	}()

	wg.Wait()

	wg.Add(2)

	go func() {

		defer func() {
			wg.Done()
		}()

		for i := 1; i <= 10; i++ {
			value, exists := m.Load(i)
			t.Logf("key: %d, exists: %v, value: %d", i, exists, value)
		}
	}()

	go func() {

		defer func() {
			wg.Done()
		}()

		for i := 11; i <= 20; i++ {
			value, exists := m.Load(i)
			t.Logf("key: %d, exists: %v, value: %d", i, exists, value)
		}

	}()

	wg.Wait()

}

func TestMyMapStoreAndLoad(t *testing.T) {
	m := NewMyMap()

	timer1 := time.AfterFunc(time.Second, func() {
		actual, exist := m.LoadAndDelete(1)
		t.Logf("delete: actual: %v, exist: %v", actual, exist)
	})

	timer2 := time.AfterFunc(time.Second, func() {
		actual, exist := m.LoadOrStore(1, 2)
		t.Logf("LoadAndStore: actual: %v, exist: %v", actual, exist)
	})

	timer3 := time.AfterFunc(time.Second, func() {
		actual, exist := m.LoadOrStore(1, 3)
		t.Logf("LoadAndStore: actual: %v, exist: %v", actual, exist)
	})

	defer func() {
		timer1.Stop()
		timer2.Stop()
		timer3.Stop()
	}()

	time.Sleep(time.Second * 5)
}

func BenchmarkMyMap(b *testing.B) {
	m := NewMyMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Store(1, 1)
			_, _ = m.Load(1)
			m.LoadOrStore(1, 1)
			m.LoadOrStore(2, 2)
			m.LoadAndDelete(2)
			m.LoadAndDelete(1)
			m.LoadAndDelete(1)
		}
	})
}
func BenchmarkSyncMap(b *testing.B) {
	m := sync.Map{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Store(1, 1)
			_, _ = m.Load(1)
			m.LoadOrStore(1, 1)
			m.LoadOrStore(2, 2)
			m.LoadAndDelete(2)
			m.LoadAndDelete(1)
			m.LoadAndDelete(1)
		}
	})
}
