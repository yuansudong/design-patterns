package iterator

type Iterator interface {
	HasNext() bool
	GetNext() interface{}
}

type Collection interface {
	GetIterator() Iterator
}

type List struct {
	index int
	sets  []interface{}
}

func New() *List {
	return &List{}
}

func (u *List) HasNext() bool {
	return u.index < len(u.sets)
}

func (u *List) GetIterator() Iterator {
	l := *u
	return &l
}

func (u *List) Add(ele interface{}) *List {
	u.sets = append(u.sets, ele)
	return u
}

func (u *List) GetNext() interface{} {
	if u.HasNext() {
		ele := u.sets[u.index]
		u.index++
		return ele
	}
	return nil
}
