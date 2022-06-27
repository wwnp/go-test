package main

import "fmt"

type Foo interface {
	getWidth() int
}

type TestStruct struct {
	width int
}

func (t TestStruct) getWidth() int {
	return t.width
}

func getWidthThroughInterface(f Foo) {
	fmt.Println(f.getWidth())
}

func main() {
	testQ := TestStruct{50}
	// fmt.Println(testQ.getWidth())
	getWidthThroughInterface(testQ)
}
