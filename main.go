package main

import "fmt"

type Shape interface {
	Area() int
}

type Rec struct {
	side int
}

func (r Rec) Area() int {
	return r.side * r.side
}

func main() {
	rec := Rec{5}
	printShape(rec)

}
func printShape(s Shape) {
	fmt.Println(s.Area())
}
