package state

import "fmt"

var (
	InOfStock  IState = new(InOfStockState)
	OutOfStock IState = new(OutOfStockState)
)

type IContext interface {
	GetNum() int
	SetNum(int)
	SetState(IState)
	Out() error
}

type IState interface {
	Out(IContext) error
}

type OutOfStockState struct {
}

func (ons *OutOfStockState) Out(ctx IContext) error {
	if ctx.GetNum() > 0 {
		ctx.SetState(InOfStock)
		return ctx.Out()
	}
	return fmt.Errorf("hub is not enough")
}

type InOfStockState struct {
}

func (ofs *InOfStockState) Out(ctx IContext) error {
	if ctx.GetNum() <= 0 {
		ctx.SetState(OutOfStock)
		return ctx.Out()
	}
	ctx.SetNum(ctx.GetNum() - 1)
	return nil
}

type Hub struct {
	st  IState
	num int
}

func NewHub() *Hub {
	return &Hub{
		st:  InOfStock,
		num: 1,
	}
}

// GetNum 获取数量
func (u *Hub) GetNum() int {
	return u.num
}

func (u *Hub) SetNum(n int) {
	u.num = n
}

func (u *Hub) Out() error {
	return u.st.Out(u)
}

func (u *Hub) SetState(state IState) {
	u.st = state
}
