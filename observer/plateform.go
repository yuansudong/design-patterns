package observer

import (
	"context"
	"fmt"
	"sync"
)

type Reader struct {
	name  string
	queue chan string
}

func NewReader(name string) *Reader {
	return &Reader{
		name:  name,
		queue: make(chan string),
	}
}

func (r *Reader) Send(msg string) {
	r.queue <- msg
}

func (r *Reader) Online(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-r.queue:
			fmt.Println(r.name, " 收取到了文章:", msg)
		}
	}
}

type Topic struct {
	sets []*Reader
}

func NewTopic() *Topic {
	return &Topic{
		sets: make([]*Reader, 0, 64),
	}
}

func (t *Topic) Join(r *Reader) {
	t.sets = append(t.sets, r)
}

func (t *Topic) Publish(msg string) {
	for _, item := range t.sets {
		item.Send(msg)
	}
}

type Plateform struct {
	topics map[string]*Topic
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewPlatefrom() *Plateform {
	inst := new(Plateform)
	inst.topics = make(map[string]*Topic)
	inst.ctx, inst.cancel = context.WithCancel(context.Background())
	return inst
}

func (p *Plateform) Exit() {
	p.cancel()
	p.wg.Wait()
}

func (p *Plateform) Publish(topic, msg string) {
	if toc, ok := p.topics[topic]; ok {
		toc.Publish(msg)
	}
}

// Subscribe 订阅
func (p *Plateform) Subscribe(topic string, name string) {
	defer p.wg.Add(1)
	r := NewReader(name)
	go func() {
		defer p.wg.Done()
		r.Online(p.ctx)
	}()
	toc, ok := p.topics[topic]
	if !ok {
		toc = NewTopic()
		p.topics[topic] = toc
	}
	toc.Join(r)
}
