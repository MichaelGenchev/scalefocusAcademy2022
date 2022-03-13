package main

import (
	"fmt"
	"math"
)



type Square struct {
	a float64
}
type Circle struct {
	radius float64
}

func NewSquare(a float64) Square {
	return Square{a: a}
}

func NewCircle(radius float64) Circle {
	return Circle{radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
	
}

func (s Square) Area() float64 {
	return s.a * s.a
}


type Shape interface {
	Area() float64
}

type Shapes []Shape

func (s Shapes) LargestArea() float64 {
	max := -99.99
	for _, shape := range s {
		if max < shape.Area() {
			max = shape.Area()
		}
	}
	return max
}

func main() {
	var shapes Shapes

	shapes = append(shapes, NewCircle(1))
	shapes = append(shapes, NewCircle(2))
	shapes = append(shapes, NewCircle(3))
	shapes = append(shapes, NewCircle(4))
	shapes = append(shapes, NewSquare(20))


	result := shapes.LargestArea()
	fmt.Println(result)

}