package areacalc

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	width  float64
	height float64
	name   string
}

func NewRectangle(width, height float64, name string) *Rectangle {
	return &Rectangle{width: width, height: height, name: name}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Type() string {
	return r.name
}

type Circle struct {
	radius float64
	name   string
}

func NewCircle(radius float64, name string) *Circle {
	return &Circle{radius: radius, name: name}
}

func (c *Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func (c *Circle) Type() string {
	return c.name
}

func AreaCalculator(figures []Shape) (string, float64) {
	var totalArea float64
	var names string

	for i, figure := range figures {
		if i > 0 {
			names += "-"
		}
		names += figure.Type()
		totalArea += figure.Area()
	}

	return names, totalArea
}
