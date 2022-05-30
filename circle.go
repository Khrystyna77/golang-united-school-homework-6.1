package golang_united_school_homework

import "math"

// Circle must satisfy to Shape
type Circle struct {
	Radius float64
}

//mycode

func (ac Circle) CalcArea() float64 {
	return math.Pi * ac.Radius * ac.Radius
}

func (pc Circle) CalcPerimeter() float64 {
	return math.Pi * (pc.Radius + pc.Radius)
}
