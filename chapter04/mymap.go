package chapter04

import "sync"

type MyMap struct {
	store map[interface{}]interface{}

	mux sync.RWMutex
}

func NewMyMap() *MyMap {
	return &MyMap{
		store: map[interface{}]interface{}{},
	}
}

func (m *MyMap) Load(key interface{}) (value interface{}, ok bool) {
	m.mux.RLock()
	defer func() {
		m.mux.RUnlock()
	}()

	value, ok = m.store[key]
	return
}

func (m *MyMap) Store(key, value interface{}) {
	m.mux.Lock()
	defer func() {
		m.mux.Unlock()
	}()

	m.store[key] = value
}

func (m *MyMap) Delete(key interface{}) {
	m.mux.Lock()
	defer func() {
		m.mux.Unlock()
	}()

	delete(m.store, key)
}

// LoadOrStore 如果存在, 返回现有值; 否则, 存储并返回给定值. 如果值存在则返回true, 如果存储了给定值, 则返回false.
func (m *MyMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.mux.RLock()

	if actual, loaded = m.store[key]; loaded {
		m.mux.RUnlock()
		return
	}

	m.mux.RUnlock()
	m.mux.Lock()
	defer func() {
		m.mux.Unlock()
	}()

	m.store[key] = value
	actual, loaded = value, false

	return
}

// LoadAndDelete 删除键的值, 并返回键的值(如果存在). 如果键值存在, 则返回true, 否则返回false
func (m *MyMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	m.mux.RLock()

	if v, exist := m.store[key]; !exist {
		m.mux.RUnlock()
		return
	} else {
		value = v
	}

	m.mux.RUnlock()
	m.mux.Lock()
	defer func() {
		m.mux.Unlock()
	}()

	delete(m.store, key)
	loaded = true
	return
}
