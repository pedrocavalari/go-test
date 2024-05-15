package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Heigth float64
	Width  float64
}

type Circle struct {
	Rad float64
}

func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Rad * c.Rad)
}
