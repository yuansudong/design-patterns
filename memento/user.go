package memento

type User struct {
	Name    string
	Age     uint8
	Address string
}

// Snap 快照
func (e *User) Snap() *UserSnapshot {
	return &UserSnapshot{
		Name:    e.Name,
		Age:     e.Age,
		Address: e.Address,
	}
}
