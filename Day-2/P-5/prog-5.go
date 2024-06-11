// 5. Using interface find the area of a triangle
package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	triangle := Triangle{Base: 5, Height: 10}

	var shape Shape = triangle

	fmt.Printf("The area of the triangle is: %.2f\n", shape.Area())
}
