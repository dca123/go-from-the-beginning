package main

import "fmt"

type Point struct {
	x int
	y int
}

type Shape interface {
	area() int
	location() Point
}

type Rectangle struct {
	x int
	y int
}

func (r Rectangle) area() int {
	return r.x * r.y
}

func (r Rectangle) location() Point {
	return Point{x: r.x, y: r.y}
}

func printArea(shape Shape) {
	fmt.Println(shape.area())
}
func main() {
	var rectangle Shape = Rectangle{x: 5, y: 2}
	printArea(rectangle)
}
