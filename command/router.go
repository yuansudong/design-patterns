package command

import (
	"fmt"
	"log"
)

const (
	defaultSlot Command = 512
)

type Command uint16

const (
	ActionHello Command = iota + 1
	ActionPing
	ActionLogin
)

type Handler interface {
	GetCmd() Command
	Execute() error
}

type PingHandler struct {
	cmd Command
}

func NewPingHandler() *PingHandler {
	return &PingHandler{cmd: ActionPing}
}
func (p *PingHandler) GetCmd() Command {
	return p.cmd
}

func (p *PingHandler) Execute() error {
	fmt.Println("execute ping")
	return nil
}

type HelloHandler struct {
	cmd Command
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{cmd: ActionHello}
}

func (p *HelloHandler) GetCmd() Command {
	return p.cmd
}

func (p *HelloHandler) Execute() error {
	fmt.Println("execute hello")
	return nil
}

type LoginHandler struct {
	cmd Command
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{cmd: ActionLogin}
}
func (p *LoginHandler) GetCmd() Command {
	return p.cmd
}

func (p *LoginHandler) Execute() error {
	fmt.Println("execute login")
	return nil
}

type Router struct {
	sets []map[Command]Handler
}

func New() *Router {
	inst := &Router{sets: make([]map[Command]Handler, defaultSlot)}
	for index := Command(0); index < defaultSlot; index++ {
		inst.sets[index] = make(map[Command]Handler)
	}
	return inst
}
func (r *Router) AddHandler(h Handler) *Router {

	cmd := h.GetCmd()
	slots := r.sets[cmd%defaultSlot]
	if _, ok := slots[cmd]; ok {
		log.Fatalln("This is repeated cmd!", cmd)
	}
	slots[cmd] = h
	return r
}

func (r *Router) GetHandler(cmd Command) Handler {

	slots := r.sets[cmd%defaultSlot]
	if handler, ok := slots[cmd]; ok {
		return handler
	}
	return nil
}
