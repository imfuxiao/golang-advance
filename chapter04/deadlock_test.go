package chapter04

import "testing"

func TestDeadlock(t *testing.T) {
	var taskMap = &MyMap2{
		m: map[int]task{},
	}
	taskMap.DoMyJob(1)
}
