package main

import (
	"fmt"
	"math"
)

// ------------------------------ TYPE STRUCT ------------------------------
// Rectangle type
type Rectangle struct {
	//Parameters in golang are called as fields
	x, y float64
}

// Circle type
type Circle struct {
	//Parameters in golang are called as fields
	x, y, r float64
}

// ------------------------------ METHOD ------------------------------
//Define a new method with receiver called area
//The purpose of area is to avoid confusion

//We want to calculate area of rectangle and a circle
//In a primitive way, we could make a separate function which is named after
//what it does, i.e, areaRectangle and areaCircle.

//In golang way, we could just attach those function with a receiver to make a method
//Receiver area() ensure that we only need to call area() to calculate its area rather
//than creating separate function with the same basic purpose (calculate area)
func (r *Rectangle) area() float64 {
	//We can access fields by using dot operator
	return r.x * r.y
}

func (c *Circle) area() float64 {

	return math.Phi * c.r * c.r
}

func main() {
	r := Rectangle{4, 5}
	c := Circle{0, 0, 9}
	fmt.Println(r.area(), c.area())
	fmt.Println(totalArea(&r, &c))
}

// ------------------------------ INTERFACE ------------------------------
// Interface allows us to compile same methods
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
