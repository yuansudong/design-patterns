package bridge

import "fmt"

type CThingBridge struct {
	color IColor
	form  IFormation
}

// Combination 组合
func (ctb *CThingBridge) Combination(c IColor, f IFormation) *CThingBridge {
	ctb.color = c
	ctb.form = f
	return ctb
}

func (ctb *CThingBridge) Description() {
	if ctb.color == nil || ctb.form == nil {
		return
	}
	fmt.Println(ctb.color.GetColor(), ctb.form.GetFormation())
}
