package mediator

import (
	"context"
	"sync"
)

type Handler func(string)

type Person struct {
	fn    Handler
	queue chan string
	id    string
}

func (p *Person) Online(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-p.queue:
			p.fn(msg)
		}
	}

}
func (p *Person) Send(msg string) {
	p.queue <- msg
}

type Forward struct {
	sets   map[string]*Person
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func New() *Forward {
	inst := &Forward{
		sets: map[string]*Person{},
	}
	inst.ctx, inst.cancel = context.WithCancel(context.Background())
	return inst
}

// Spawn 实例化一个接口
func (f *Forward) Spawn(address string, fn Handler) {
	p := &Person{
		fn:    fn,
		queue: make(chan string, 1),
		id:    address,
	}
	f.sets[address] = p
	go func() {
		f.wg.Add(1)
		defer f.wg.Done()
		p.Online(f.ctx)
	}()
}

func (f *Forward) Exit() {
	f.cancel()
	f.wg.Wait()
}

func (f *Forward) Send(address, msg string) {
	if top, ok := f.sets[address]; ok {
		top.Send(msg)
	}
}
