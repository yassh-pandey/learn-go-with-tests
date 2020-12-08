package compositetypes

import "math"

// Shape interface defines a generic geometric shape that has an area function
type Shape interface {
	Area() float64
}

// Rectangle defines a generic rectangle with a given width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter method for type Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area method for type Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle type defines a circle with a given radius
type Circle struct {
	Radius float64
}

// Perimeter method for type Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area method for type Circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Triangle is a closed polygon with three sides
type Triangle struct {
	Base   float64
	Height float64
}

// Area method for shape triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
