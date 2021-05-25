package visitor

import "fmt"

type Visitor interface {
	ForSquare(*Square)
	ForCircle(*Circle)
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) ForSquare(s *Square) {
	fmt.Println("计算正方形的面积的具体计算公式")
}

func (a *MiddleCoordinates) ForCircle(c *Circle) {
	fmt.Println("计算圆的面积的具体计算公式")
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.ForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.ForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}
