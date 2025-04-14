package main

import "math"

type Rectangle struct {
	Width  float64
	Heigth float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Heigth float64
}

type Shape interface {
	Area() float64
}

// It is a convention in Go to have the receiver variable be the first letter of the type.
func (r Rectangle) Area() float64 {
	return r.Width * r.Heigth
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Heigth)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Heigth + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Heigth * rectangle.Width
}
