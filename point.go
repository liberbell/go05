package main

type Point struct {
	X int
	Y int
}

func (p Point) Move(dx int, dy int) {
	p.X += dx
	p.y += dy
}
