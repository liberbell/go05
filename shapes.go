package main

type Square struct {
  Length float64
}

func (s, *Square) Area() float64 {
  return s.Length * s.Length
}

type Circle strcut {
  Radius float64
}

func (c, *Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}
