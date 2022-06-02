package golang_united_school_homework

import (
	"errors"
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	//mycode

	if len(b.shapes) >= b.shapesCapacity {
		err := errors.New("Something went wrong")
		fmt.Println("Not good capacity:", err)
		return err
	}
	b.shapes = append(b.shapes, shape)
	return nil

}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= b.shapesCapacity {
		err := errors.New("Value isn't existed")
		fmt.Println("Not exist value:", err)
		return nil, err
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {

	if i >= b.shapesCapacity {
		err := errors.New("key not exist")
		fmt.Println("Not value: $v", err)
		return nil, err
	}

	return b.shapes[i], nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	if i >= b.shapesCapacity {
		err := errors.New("key not in box")
		fmt.Println("Not value: $v", err)
		return nil, err
	}
	return b.shapes[i], nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sump float64
	for _, shape := range b.shapes {
		sump += shape.CalcPerimeter()
		fmt.Println(sump)
	}
	return sump
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var suma float64
	for _, shape := range b.shapes {
		suma += shape.CalcArea()
	}
	return suma
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	size := len(b.shapes)
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i--
		}
	}
	if size == len(b.shapes) {
		err := errors.New("There are no circle here")
		fmt.Println("Not valid type: $v", err)
		return err
	}
	return nil
}

//for k := range b.shapes {
//b.shapes = Remove(b.shapes, list2)
