package main

type Shape interface {
	getArea() float64
}

type Rectangle struct {
	width  float64
	height float64
}	

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.height
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}


func printArea(s Shape) {
	area := s.getArea()
	println("Area:", area)
}	


