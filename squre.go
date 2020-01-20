package main

type Squre struct {
	Center Point
	Length int
}

func NewSqure(x int, y int, length int) (*Squre error) {
  if length <= 0 {
    return nil, fmt.Errorf("length must be > 0 (was %d)\n", length)
  }

  s := &Squre{
    Center: Point{x, y},
    Length: length,
  }
}
