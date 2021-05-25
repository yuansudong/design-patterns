package factory

import "errors"

// GetPerson 用于获取Person数据
func GetPerson(c Category) (IPerson, error) {
	switch c {
	case C_STUDENT:
		return new(student), nil
	case C_TEACHER:
		return new(teacher), nil
	}
	return nil, errors.New("Unknown Category")
}
