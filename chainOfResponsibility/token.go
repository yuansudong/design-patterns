package chainOfResponsibility

import "fmt"

type Token struct {
	next Handler
}

func (t *Token) SetNext(h Handler) {
	t.next = h
}
func (t *Token) Execute(ctx IContext) (err error) {
	if ctx.GetToken() != "test" {
		return fmt.Errorf("token is illege!")
	}
	if t.next != nil {
		err = t.next.Execute(ctx)
	}
	return err
}
