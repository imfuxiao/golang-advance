package chapter04

import (
	"context"
)

type MyContext struct {
	ctx context.Context
}

func NewMyContext(ctx context.Context) *MyContext {
	return &MyContext{
		ctx: ctx,
	}
}

func (m *MyContext) WithValue(parent context.Context, key, val interface{}) context.Context {
	m.ctx = context.WithValue(parent, key, val)
	return m.ctx
}

func (m *MyContext) Value(key interface{}) interface{} {
	return m.ctx.Value(key)
}
