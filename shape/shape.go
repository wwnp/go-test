package shape

import "math"

type Shape interface {
	getArea() float64
	// getPerimetr() float64
}

type Circle struct {
	radius float64
}

func NewCircle(length float64) Circle {
	return Circle{
		radius: length,
	}
}

func (c Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
