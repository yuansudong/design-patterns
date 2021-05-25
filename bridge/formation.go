package bridge

type IFormation interface {
	GetFormation() string
}

type CTreeFormation struct {
}

func (ctf *CTreeFormation) GetFormation() string {

	return "树"
}

type CFlowerFormation struct {
}

func (cff *CFlowerFormation) GetFormation() string {
	return "花"
}
