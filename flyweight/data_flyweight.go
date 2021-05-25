package flyweight

import "fmt"

type StaticData struct {
	speed  float64
	sprite string
	harma  float64
}

func NewStaticData() *StaticData {
	return &StaticData{
		speed:  1.0,
		sprite: "./ac.png",
		harma:  666,
	}
}

func (sd *StaticData) GetSprite() string {

	return sd.sprite
}

func (sd *StaticData) GetSpeed() float64 {
	return sd.speed
}

// GetHarma
func (sd *StaticData) GetHarma() float64 {
	return sd.harma
}

type DynamicData struct {
	position string
	staticed *StaticData
}

//
func (dd *DynamicData) SetStatic(sd *StaticData) *DynamicData {
	dd.staticed = sd
	return dd
}

func (dd *DynamicData) Show() {
	fmt.Printf("%+v\n", dd)
}

// SetPosition 用于设置位置
func (dd *DynamicData) SetPosition(pos string) *DynamicData {
	dd.position = pos
	return dd
}
