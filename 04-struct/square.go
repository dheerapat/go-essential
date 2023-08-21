package main

import (
	"fmt"
	"log"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length == 0 {
		return nil, fmt.Errorf("square must not be 0 in length")
	}

	s := Square{
		X:      x,
		Y:      y,
		Length: length,
	}

	return &s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

func (s Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
