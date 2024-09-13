package classes

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
	Color  *Color
}

func NewCircle(radius float64, color *Color) *Circle {
	return &Circle{Radius: radius, Color: color}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f, color=%s, area=%.2f)", c.Radius, c.Color, c.Area())
}
