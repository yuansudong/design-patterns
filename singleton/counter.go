package singleton

import (
	"sync"
	"sync/atomic"
)

var (
	inst *Counter
	once sync.Once
)

func init() {
	once.Do(func() {
		inst = new(Counter)
	})
}

func GetInstance() *Counter {
	return inst
}

type Counter struct {
	count int64
}

func (c *Counter) Incr() (int64, error) {
	iNew := atomic.AddInt64(&c.count, 1)
	return iNew, nil
}

func (c *Counter) Dec() (int64, error) {
	iNew := atomic.AddInt64(&c.count, -1)
	return iNew, nil
}

func (c *Counter) Get() (int64, error) {
	iVal := atomic.LoadInt64(&c.count)
	return iVal, nil
}
