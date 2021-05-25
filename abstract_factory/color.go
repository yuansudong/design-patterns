package abstract_factory

type EColor int32

const (
	ECO_UNKNOWN EColor = 0
	ECO_RED     EColor = 1
	ECO_BLUE    EColor = 2
)

type IColor interface {
	GetColor() (string, error)
}

type CRedColor struct {
}

func (c *CRedColor) GetColor() (string, error) {
	return "RED", nil
}

//
type CBlueColor struct {
}

func (c *CBlueColor) GetColor() (string, error) {
	return "BLUE", nil
}
