package main

import "fmt"

type Shape interface {
	ShapeArea
	ShapePerimeter
}
type ShapeArea interface {
	getArea() int
}
type ShapePerimeter interface {
	getPerimeter() int
}

type Rectangle struct {
	width int
}

func (r Rectangle) getArea() int {
	return r.width
}
func (r Rectangle) getPerimeter() int {
	return 123
}

func getArea(s Shape) {
	fmt.Println(s.getArea())
}
func getPerimeter(s Shape) {
	fmt.Println(s.getPerimeter())
}

func main() {
	testQ := Rectangle{50}
	// fmt.Println(testQ.getArea())
	getArea(testQ)
	getPerimeter(testQ)
}
