package chapter04

import "sync"

type MyLock struct {
	c chan int8

	mux sync.Mutex
}

func NewMyLock() MyLock {
	return MyLock{
		c: make(chan int8, 1),
	}
}

func (m *MyLock) TryLock() bool {
	if len(m.c) > 0 {
		return false
	}
	m.Lock()
	return true
}

func (m *MyLock) Lock() {
	select {}
	m.mux.Lock()
	m.c <- 1
}

func (m *MyLock) Unlock() {
	m.mux.Unlock()
	<-m.c
}
