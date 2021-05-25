package builder

type Type int32

const (
	MYSQL Type = iota + 1
)

func Get(t Type) IQueryBuilder {
	switch t {
	case MYSQL:
		return new(CMysqlBuilder)
	}
	return nil
}
