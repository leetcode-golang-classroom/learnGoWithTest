package structs_methods_interfaces

import "math"

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
