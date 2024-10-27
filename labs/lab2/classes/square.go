package classes

import "fmt"

type Square struct {
	Side  float64
	Color *Color
}

func NewSquare(side float64, color *Color) *Square {
	return &Square{Side: side, Color: color}
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) String() string {
	return fmt.Sprintf("Square(side=%.2f, color=%s, area=%.2f)", s.Side, s.Color, s.Area())
}
