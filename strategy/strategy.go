package state

import "fmt"

var (
	Offline IState = new(OfflineState)
	Online  IState = new(OnlineState)
)

type IContext interface {
}

type IState interface {
	Send(IContext, string)
}

type OnlineState struct {
}

func (ons *OnlineState) Send(ctx IContext, msg string) {
	fmt.Println("我在线,消息缓存后,直接下发:", msg)
}

type OfflineState struct {
}

func (ofs *OfflineState) Send(ctx IContext, msg string) {
	fmt.Println("我不在线,将消息存储到离线消息后.等待用户上线获取:", msg)
}

type User struct {
	st IState
}

func NewUser() *User {
	return &User{
		st: Offline,
	}
}

func (u *User) Send(msg string) {
	u.st.Send(u, msg)
}

func (u *User) SwitchStateTo(state IState) {
	u.st = state
}
