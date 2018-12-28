package gotour

import (
  "math"
)

type Shape interface {
  Area() float64
  Circumference() float64
}

type Circle struct {
  Radius float64
}

type Rectangle struct {
  Breadth,Width float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
  return r.Breadth * r.Width
}

func (c Circle) Circumference() float64{
  return 2*math.Pi*c.Radius
}

func (r Rectangle) Circumference() float64{
  return 2 *  (r.Breadth + r.Width)
}
