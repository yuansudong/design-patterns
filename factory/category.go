package factory

type Category int32

const (
	C_UNKNOWN = 0
	C_TEACHER = 1
	C_STUDENT = 2
)

var enumSet map[Category]string = map[Category]string{
	C_UNKNOWN: "C_UNKNOWN",
	C_TEACHER: "C_TEACHER",
	C_STUDENT: "C_STUDENT",
}

func (C Category) String() string {
	if str, ok := enumSet[C]; ok {
		return str
	}
	return ""
}
