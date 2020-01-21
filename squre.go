package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (was %d)\n", length)
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("Error: can`t create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
