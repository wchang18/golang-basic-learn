package lesson2

import "testing"

type Shape interface {
	GetArea() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) GetArea() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) GetArea() float64 {
	return r.Width * r.Height
}

func CalculateArea(shape Shape) float64 {
	return shape.GetArea()
}

func TestCalculateArea(t *testing.T) {
	circle := &Circle{Radius: 5}
	area := CalculateArea(circle)
	t.Log(area)
	rectangle := &Rectangle{Width: 4, Height: 6}
	area = CalculateArea(rectangle)
	t.Log(area)
}
