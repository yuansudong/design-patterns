package abstract_factory

type EClothing int32

const (
	ECL_UNKNOWN EClothing = 0
	ECL_T_SHIRT EClothing = 1
	ECL_SHOES   EClothing = 2
)

type IClothing interface {
	GetClothing() (string, error)
}

type CTShirtClothing struct {
}

func (c *CTShirtClothing) GetClothing() (string, error) {

	return "T-SHIRT", nil
}

//
type CShoesClothing struct {
}

func (c *CShoesClothing) GetClothing() (string, error) {
	return "SHOES", nil
}
