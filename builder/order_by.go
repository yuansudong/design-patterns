package builder

import "strings"

type IOrderBy interface {
	Columns(string) IOrderBy
	ASC() IOrderBy
	Desc() IOrderBy
	Build() (string, error)
}
type COrderBy struct {
	columns string
	order   bool
}

func (c *COrderBy) Columns(string) IOrderBy {
	return c
}
func (c *COrderBy) ASC() IOrderBy {
	return c
}
func (c *COrderBy) Desc() IOrderBy {
	c.order = true
	return c
}
func (c *COrderBy) Build() (string, error) {
	v := "ASC"
	if c.order {
		v = "DESC"
	}
	return strings.Join([]string{c.columns, v}, " "), nil
}
