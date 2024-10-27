package classes

type Color struct {
	Name string
}

func NewColor(name string) *Color {
	return &Color{Name: name}
}

func (c *Color) String() string {
	return c.Name
}
