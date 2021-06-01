package chapter04

import "sync"

type task struct{}

type MyMap2 struct {
	m   map[int]task
	mux sync.RWMutex
}

func (m *MyMap2) finishJob(t task, id int) {
	m.mux.Lock()
	defer m.mux.Unlock()

	// finish task
	delete(m.m, id)
}

func (m *MyMap2) DoMyJob(taskID int) {
	m.mux.RLock()
	// defer m.mux.RUnlock()

	t := m.m[taskID]
	m.mux.RUnlock()

	// 这里有问题: 调用函数, 函数内部又加了锁, 导致重复上锁, 产生异常
	// 修改上面的defer m.mux.RUnlock()
	// 在读取map值后, 直接释放读锁
	m.finishJob(t, taskID)
}
