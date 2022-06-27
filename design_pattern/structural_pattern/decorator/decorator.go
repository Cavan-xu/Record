package decorator

/*
	装饰器模式：主要解决继承关系过于复杂的问题，通过组合来代替继承，主要作用是给原始类增加功能
*/

type Draw interface {
	Draw() string
}

type Square struct{}

func (s *Square) Draw() string {
	return "draw square"
}

type ColorSquare struct {
	square Square
	color  string
}

func (c *ColorSquare) Draw() string {
	return c.square.Draw() + " color is: " + c.color
}

func NewColorSquare(square Square, color string) *ColorSquare {
	return &ColorSquare{
		square: square,
		color:  color,
	}
}
