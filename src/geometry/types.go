package main

import "math"

type Shape struct {
	height int
	width  int
}

type Vector struct {
	x, y float64
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f))
}

func (v Vector) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
