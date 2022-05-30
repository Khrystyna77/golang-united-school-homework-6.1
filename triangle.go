package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape
type Triangle struct {
	Side float64
}

//mycode

func (tr *Triangle) CalcArea() float64 {
	s := (tr.Side + tr.Side + tr.Side) / 2
	return math.Sqrt(s * (s - tr.Side) * (s - tr.Side) * (s - tr.Side))
}

func (trp *Triangle) CalcPerimeter() float64 {
	return trp.Side + trp.Side + trp.Side
}
