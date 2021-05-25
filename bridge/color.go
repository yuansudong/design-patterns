package bridge

type IColor interface {
	GetColor() string
}

type CRedColor struct {
}

func (crc *CRedColor) GetColor() string {
	return "红色的"
}

type CBlueColor struct {
}

func (cbc *CBlueColor) GetColor() string {
	return "蓝色的"
}
