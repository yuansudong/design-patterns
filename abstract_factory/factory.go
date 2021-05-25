package abstract_factory

import "errors"

type IFactory interface {
	CreateClothing(EClothing) (IClothing, error)
	CreateColor(EColor) (IColor, error)
}

func GetFactory() IFactory {

	return new(CChoiceFactory)
}

type CChoiceFactory struct {
}

// CreateClothing 创建服装
func (ccf *CChoiceFactory) CreateClothing(e EClothing) (IClothing, error) {
	switch e {
	case ECL_SHOES:
		return new(CShoesClothing), nil
	case ECL_T_SHIRT:
		return new(CTShirtClothing), nil
	}
	return nil, errors.New("Unknown EClothing")
}

// CreateColor 创建颜色
func (ccf *CChoiceFactory) CreateColor(e EColor) (IColor, error) {
	switch e {
	case ECO_RED:
		return new(CRedColor), nil
	case ECO_BLUE:
		return new(CBlueColor), nil
	}
	return nil, errors.New("Unknown EColor")
}
