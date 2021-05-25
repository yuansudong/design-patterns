package chainOfResponsibility

import "time"

type Bussiness struct {
	next Handler
}

func (b *Bussiness) SetNext(n Handler) {
	b.next = n
}

// Execute
func (b *Bussiness) Execute(ctx IContext) error {
	if b.next != nil {
		return b.next.Execute(ctx)
	}
	time.Sleep(200 * time.Millisecond)
	return nil
}
