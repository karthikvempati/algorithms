package gotour

import (
  "testing"
  "math"
  )


func TestArea(t *testing.T)  {
  s := Circle{Radius:1}
  value := s.Area()
  if value != math.Pi {
    t.Error("Expected math.Pi, got ", value)
  }
}

func TestCircumference(t *testing.T) {
  r := Rectangle{Breadth:1, Width: 1}
  v := r.Circumference()
  if v != 4 {
    t.Error("Expected 4, got ", v)
  }
}
