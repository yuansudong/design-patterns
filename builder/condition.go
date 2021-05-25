package builder

import "strings"

type ICondition interface {
	Columns(string) ICondition
	Value(interface{}) ICondition
	Build() (string, error)
}

type CLTOperator struct {
	columns string
	value   interface{}
}

func (clt *CLTOperator) Columns(arg string) ICondition {
	clt.columns = arg
	return clt
}
func (clt *CLTOperator) Value(arg interface{}) ICondition {
	clt.value = arg
	return clt
}
func (clt *CLTOperator) Build() (string, error) {
	v, err := sqlValue(clt.value)
	if err != nil {
		return v, err
	}
	return strings.Join([]string{clt.columns, "<", v}, " "), nil
}

type CLTEOperator struct {
	columns string
	value   interface{}
}

func (clt *CLTEOperator) Columns(arg string) ICondition {
	clt.columns = arg
	return clt
}

func (clt *CLTEOperator) Value(arg interface{}) ICondition {
	clt.value = arg
	return clt
}
func (clt *CLTEOperator) Build() (string, error) {
	v, err := sqlValue(clt.value)
	if err != nil {
		return v, err
	}
	return strings.Join([]string{clt.columns, "<=", v}, " "), nil
}

type CGTOperator struct {
	columns string
	value   interface{}
}

func (clt *CGTOperator) Columns(arg string) ICondition {
	clt.columns = arg
	return clt
}
func (clt *CGTOperator) Value(arg interface{}) ICondition {
	clt.value = arg
	return clt
}
func (clt *CGTOperator) Build() (string, error) {
	v, err := sqlValue(clt.value)
	if err != nil {
		return v, err
	}
	return strings.Join([]string{clt.columns, ">", v}, " "), nil
}

type CGTEOperator struct {
	columns string
	value   interface{}
}

func (clt *CGTEOperator) Columns(arg string) ICondition {
	clt.columns = arg
	return clt
}
func (clt *CGTEOperator) Value(arg interface{}) ICondition {
	clt.value = arg
	return clt
}
func (clt *CGTEOperator) Build() (string, error) {
	v, err := sqlValue(clt.value)
	if err != nil {
		return v, err
	}
	return strings.Join([]string{clt.columns, ">=", v}, " "), nil
}

type CEQOperator struct {
	columns string
	value   interface{}
}

func (clt *CEQOperator) Columns(arg string) ICondition {
	clt.columns = arg
	return clt
}
func (clt *CEQOperator) Value(arg interface{}) ICondition {
	clt.value = arg
	return clt
}
func (clt *CEQOperator) Build() (string, error) {
	v, err := sqlValue(clt.value)
	if err != nil {
		return v, err
	}
	return strings.Join([]string{clt.columns, "=", v}, " "), nil
}

type CNEQOperator struct {
	columns string
	value   interface{}
}

func (clt *CNEQOperator) Columns(arg string) ICondition {
	clt.columns = arg
	return clt
}
func (clt *CNEQOperator) Value(arg interface{}) ICondition {
	clt.value = arg
	return clt
}
func (clt *CNEQOperator) Build() (string, error) {
	v, err := sqlValue(clt.value)
	if err != nil {
		return v, err
	}
	return strings.Join([]string{clt.columns, "!=", v}, " "), nil
}
