package chapter04

import (
	"context"
	"testing"
)

func TestMyContext(t *testing.T) {
	ctx := context.Background()
	c := &MyContext{}

	c2 := c.WithValue(ctx, 1, 1)
	t.Logf("c2: key = %v, value = %v", 1, c2.Value(1))

	t.Logf("ctx: key = %v, value = %v", 1, c.Value(1))
}
