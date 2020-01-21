package main

import "math"

type Square struct {
	Length float64
}

func (s, *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c, *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
