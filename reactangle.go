package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

//mycode

func (rec Rectangle) CalcArea() float64 {
	return rec.Weight * rec.Height
}

func (recp Rectangle) CalcPerimeter() float64 {
	return 2.0 * (recp.Weight + recp.Height)
}
