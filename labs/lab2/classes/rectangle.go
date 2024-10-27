package classes

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
	Color  *Color
}

func NewRectangle(width, height float64, color *Color) *Rectangle {
	return &Rectangle{Width: width, Height: height, Color: color}
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle(width=%.2f, height=%.2f, color=%s, area=%.2f)", r.Width, r.Height, r.Color, r.Area())
}
