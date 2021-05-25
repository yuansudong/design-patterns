package template

import (
	"fmt"
)

type IDataCleaning interface {
	Load()
	Filter()
	Clean()
	Show()
}

type BaseDataClean struct {
	sets []int
}

func (bdc *BaseDataClean) Load() {

}
func (bdc *BaseDataClean) Show() {
	fmt.Println("当前的数据是:", bdc.sets)
}
func (bdc *BaseDataClean) Clean() {
	fmt.Println("数据已经清理")
}
func (bdc *BaseDataClean) Filter() {
	fmt.Println("数据已经过滤")
}

type XmlDataClean struct {
	BaseDataClean
}

func (xdc *XmlDataClean) Load() {
	xdc.sets = []int{1, 2, 3, 4, 5}
}

type CsvDataClean struct {
	BaseDataClean
}

func (xdc *CsvDataClean) Load() {
	xdc.sets = []int{6, 7, 8}
}

type Control struct {
	driver IDataCleaning
}

func (c *Control) SetDriver(i IDataCleaning) {
	c.driver = i
}
func (c *Control) Execute() {
	c.driver.Load()
	c.driver.Filter()
	c.driver.Show()
	c.driver.Clean()
}
