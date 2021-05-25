package prototype

import (
	"fmt"
	"time"
)

type NPC struct {
	PersistentProperties1 int
	PersistentProperties2 int
	PersistentProperties3 int
	DynamicProperties1    int
	DynamicProperties2    int
	DynamicProperties3    int
}

func New() *NPC {
	inst := new(NPC)
	// 模拟PersistentProperties1 加载时间
	time.Sleep(time.Second)
	inst.PersistentProperties1 = 1
	// 模拟PersistentProperties2 加载时间
	time.Sleep(time.Second)
	inst.PersistentProperties2 = 2
	// 模拟PersistentProperties3 加载时间
	time.Sleep(time.Second)
	inst.PersistentProperties3 = 3
	return inst
}

// Init 初始化
func (n *NPC) Init(one, two, three int) *NPC {
	n.DynamicProperties1 = n.PersistentProperties1 * one
	n.DynamicProperties2 = n.PersistentProperties2 * two
	n.DynamicProperties3 = n.PersistentProperties3 * three
	return n
}
func (n *NPC) Result() {
	fmt.Println("动态属性1:", n.DynamicProperties1, "动态属性2:", n.DynamicProperties2, "动态属性3:", n.DynamicProperties3)
}

func (n *NPC) Clone() *NPC {
	reuseNpc := *n
	return &reuseNpc

}
