package composite

type BoxType int32

const (
	Apple BoxType = iota + 1 // Boxt代表是盒子
	Banana
	Pineapple
	Orange
	Grape
	Longan
	Litchi
)

type IBox interface {
	SetPrice(float64) IBox
	GetPrice() float64
	Add(IBox) IBox
}

type Box struct {
	leafs []IBox
	price float64
}

func (b *Box) Add(g IBox) IBox {
	b.leafs = append(b.leafs, g)
	return b
}

// SetPrice
func (b *Box) SetPrice(p float64) IBox {
	b.price = p
	return b
}

func (b *Box) GetPrice() float64 {
	child := 0.00
	for _, ib := range b.leafs {
		child += ib.GetPrice()
	}
	return b.price + child
}
