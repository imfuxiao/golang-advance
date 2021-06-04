package chapter04

import (
	"context"
)

type MyContext struct {
	cache map[interface{}]interface{}
}

func NewMyContext() *MyContext {
	return &MyContext{
		cache: map[interface{}]interface{}{},
	}
}

func (m *MyContext) WithValue(parent context.Context, key, val interface{}) context.Context {
	ctx := context.WithValue(parent, key, val)
	m.cache[key] = val
	return ctx
}

func (m *MyContext) Value(key interface{}) interface{} {
	return m.cache[key]
}
