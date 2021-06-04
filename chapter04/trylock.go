package chapter04

type MyLock struct {
	c chan struct{}
}

func NewMyLock() MyLock {
	return MyLock{
		c: make(chan struct{}, 1),
	}
}

func (m *MyLock) TryLock() bool {
	select {
	case m.c <- struct{}{}:
		return true
	default:
		return false
	}
}

func (m *MyLock) Lock() {
	m.c <- struct{}{}
}

func (m *MyLock) Unlock() {
	<-m.c
}
