package main

import (
	"fmt"
	"math"
	"reflect"
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
	myType := reflect.ValueOf(s)
	fmt.Println(myType)

	// myType := reflect.TypeOf(s)
	// fmt.Println(myType.Name())
	// fmt.Println(myType.Kind())
	// fieldInfo := myType.Field(0)
	// fmt.Println(fieldInfo.Name)

	// fmt.Println(myType.Elem().Name())
	// fmt.Println(myType.NumField())
	switch x := s.(type) {
	case circle:
		fmt.Printf("I'm a circle and My type is %T \n", x)
	case rectangle:
		fmt.Printf("I'm a rectangle and My type is %T \n", x)
	default:
		fmt.Println("None of the above types")
	}

	fmt.Println("_________________________________________")
}

func main() {
	r := rectangle{width: 5, height: 8}
	c := circle{radius: 6}

	// n := 3

	Calculate(r)
	Calculate(c)
	// Calculate(&n)
}
