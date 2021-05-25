package builder

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func sqlValue(arg interface{}) (string, error) {
	switch v := arg.(type) {
	case string:
		return strconv.Quote(v), nil
	case int8, int64, int32, int, int16, uint8, uint64, uint32, uint, uint16:
		return fmt.Sprint(v), nil
	}
	data, err := json.Marshal(arg)
	if err != nil {
		return "", err
	}
	return strconv.Quote(string(data)), nil
}

type IQueryBuilder interface {
	Select(...string) IQueryBuilder
	From(string) IQueryBuilder
	Where(...ICondition) IQueryBuilder
	Having(...ICondition) IQueryBuilder
	Offset(int) IQueryBuilder
	Limit(int) IQueryBuilder
	Build() (string, error)
}

type CMysqlBuilder struct {
	from    string
	columns []string
	wheres  []ICondition
	havings []ICondition
	limit   int
	offset  int
	orderby []IOrderBy
}

func (c *CMysqlBuilder) From(args string) IQueryBuilder {
	c.from = args
	return c
}
func (c *CMysqlBuilder) Select(args ...string) IQueryBuilder {
	c.columns = append(c.columns, args...)
	return c
}
func (c *CMysqlBuilder) Where(args ...ICondition) IQueryBuilder {
	c.wheres = append(c.wheres, args...)
	return c
}
func (c *CMysqlBuilder) Having(args ...ICondition) IQueryBuilder {
	c.havings = append(c.havings, args...)
	return c
}
func (c *CMysqlBuilder) Offset(offset int) IQueryBuilder {
	c.offset = offset
	return c
}
func (c *CMysqlBuilder) Limit(limit int) IQueryBuilder {
	c.limit = limit
	return c
}
func (c *CMysqlBuilder) Build() (string, error) {
	sql := []string{}
	fields := "*"
	if len(c.columns) != 0 {
		fields = strings.Join(c.columns, " ")
	}
	sql = append(sql, "SELECT", fields, "FROM", c.from)
	if len(c.wheres) != 0 {
		sql = append(sql, "WHERE")
		for _, condition := range c.wheres {
			sVal, sErr := condition.Build()
			if sErr != nil {
				return "", sErr
			}
			sql = append(sql, sVal)
		}
	}
	if len(c.havings) != 0 {
		sql = append(sql, "HAVING")
		for _, condition := range c.havings {
			sVal, sErr := condition.Build()
			if sErr != nil {
				return "", sErr
			}
			sql = append(sql, sVal)
		}
	}
	if len(c.orderby) != 0 {
		sql = append(sql, "ORDER BY")
		for _, condition := range c.havings {
			sVal, sErr := condition.Build()
			if sErr != nil {
				return "", sErr
			}
			sql = append(sql, sVal)
		}
	}
	if c.limit != 0 {
		sql = append(sql, "LIMIT", fmt.Sprint(c.limit))
	}
	if c.offset != 0 {
		sql = append(sql, "OFFSET", fmt.Sprint(c.offset))
	}

	return strings.Join(sql, " "), nil
}
