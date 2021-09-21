package main

import (
	"fmt"
	"math"
)

var i interface{}

type shape interface {
	area() float64
	perimeter() float64
}
type polygon interface {
	getSides() int
}

type rectangle struct {
	width, height float64
	sides         int
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}
func (r rectangle) getSides() int {
	return r.sides
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// func Calculate(s shape) {
func Calculate(s interface{}) {
	// type assertion in GO
	// rect := s.(rectangle)
	// fmt.Println(rect.getSides())

	// rect := s.(polygon)
	// fmt.Println(rect.getSides())
	// fmt.Printf("My type is %T \n", rect)

	// if rect, ok := s.(polygon); ok {
	// 	fmt.Println(rect.getSides())
	// 	fmt.Printf("My type is %T \n", rect)
	// }

	// fmt.Printf("My type is %T \n", s)
	// fmt.Printf("My type is %v \n", s)

	// fmt.Printf("The area of the shape is %0.2f \n", s.area())
	// fmt.Printf("The premeter of the shape is %0.2f \n", s.perimeter())

	switch x := s.(type) {
	case circle:
		fmt.Printf("My type is %T \n", x)
		fmt.Printf("The area of the shape is %0.2f \n", x.area())
		fmt.Printf("The premeter of the shape is %0.2f \n", x.perimeter())
	case rectangle:
		fmt.Printf("My type is %T \n", x)
		fmt.Printf("The area of the shape is %0.2f \n", x.area())
		fmt.Printf("The premeter of the shape is %0.2f \n", x.perimeter())
	case int:
		fmt.Printf("My type is %T \n", x)
	default:
		fmt.Println("None of the above types")
	}
}

func main() {
	r := rectangle{width: 5, height: 8}
	c := circle{radius: 6}

	Calculate(r)
	Calculate(c)

	shapes := []shape{
		rectangle{width: 5, height: 8},
		circle{radius: 6},
		circle{radius: 9.2},
		rectangle{width: 4.3, height: 11.8},
	}

	for _, v := range shapes {
		Calculate(v)
	}

	i = 20
	fmt.Printf("The value of i is %v \n", i)
	Calculate(i)

	i = "some string"
	fmt.Printf("The value of i is %v \n", i)
	Calculate(i)

	i = true
	Calculate(i)
	fmt.Printf("The value of i is %v \n", i)
}
