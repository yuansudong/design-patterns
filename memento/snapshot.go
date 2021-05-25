package memento

import (
	"encoding/json"
	"io/ioutil"
)

type UserSnapshot struct {
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
}

// Backup 备份
func (m *UserSnapshot) Backup(file string) (err error) {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0655)
}

func (m *UserSnapshot) Recover(file string) (*User, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, m); err != nil {
		return nil, err
	}
	return &User{
		Name:    m.Name,
		Age:     m.Age,
		Address: m.Address,
	}, nil
}
