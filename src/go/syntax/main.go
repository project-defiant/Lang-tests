// Run: go run src/go/main.go
// Build: go build src/go/main.go
// This is a simple program to demonstrate the basic syntax of Go programming language.
package main

import "fmt"

func main() {
	// variables declaration
	var isTrue bool = true
	fmt.Println("isTrue: ", isTrue)
	var isFalse bool = false
	fmt.Println("isFalse: ", isFalse)
	var Int int = 1
	fmt.Println("Int: ", Int)
	var Float float64 = 1.1
	fmt.Println("Float: ", Float)
	var String string = "Hello, World!"
	fmt.Println("String: ", String)
	var Rune rune = 'A'
	fmt.Println("Rune: ", Rune)
	const Pi = 3.14
	fmt.Println("Pi: ", Pi)

	// boolean operators
	if isTrue && !isFalse {
		fmt.Println("isTrue && !isFalse: ", isTrue && isFalse)
	}
	if isTrue || isFalse {
		fmt.Println("isTrue || isFalse: ", isTrue || isFalse)
	}
	if !isFalse {
		fmt.Println("!isFalse: ", !isFalse)
	}
	if isTrue != isFalse {
		fmt.Println("isTrue != isFalse: ", isTrue != isFalse)
	}

	// arithmetic operators
	Int = increment(Int, 10)
	fmt.Println("increment(Int, 10): ", Int)

	// multiple return values
	x, _ := multipleReturn(Int, 10)
	fmt.Println("multipleReturn(Int, 10): ", x)

	// named return values
	width, height := namedReturns()
	fmt.Println("namedReturns(): ", width, height)

	// structs

	var point Point = Point{1, 2}
	fmt.Println("point: ", point)

	var rectangle Rectangle = Rectangle{Point{1, 2}, Point{3, 4}, Point{5, 6}, Point{7, 8}}
	fmt.Println("rectangle: ", rectangle)

	var circle Circle = Circle{Point{1, 2}, 3}

	fmt.Println("circle: ", circle)

	// methods on structs

	fmt.Println("rectangle.area(): ", rectangle.area())

	// interfaces

	printShape(rectangle)

	// error handling

	// panic("Panic!")

}

type Shape interface {
	area() int
	perimeter() int
}

type Point struct {
	x int
	y int
}

type Rectangle struct {
	x1 Point
	x2 Point
	x3 Point
	x4 Point
}

// embedded structs
type Circle struct {
	Point
	radius int
}

func (r Rectangle) area() int {
	var l1 = r.x2.x - r.x1.x
	var l2 = r.x4.y - r.x1.y
	return l1 * l2
}

func (r Rectangle) perimeter() int {
	var l1 = r.x2.x - r.x1.x
	var l2 = r.x4.y - r.x1.y
	return 2 * (l1 + l2)
}

func printShape(s Shape) {
	msg := fmt.Sprintf("Area: %d, Perimeter: %d", s.area(), s.perimeter())
	fmt.Println(msg)
}

func increment(x int, y int) int {
	return x + y
}

func multipleReturn(x int, y int) (int, int) {
	return x + y, x - y
}

func namedReturns() (width int, height int) {
	width = 1
	height = 2
	return width, height
}
