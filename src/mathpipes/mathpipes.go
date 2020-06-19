package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Op struct {
	x float64
	y float64
	f func(float64, float64) float64
}

var c chan Op = make(chan Op)

func fmult(x, y float64) float64 {
	return x * y
}

func fsum(x, y float64) float64 {
	return x + y
}

func send(op Op) {
	c <- op
}

func execute() {
	op := <-c
	res := op.f(op.x, op.y)
	fmt.Printf("%v = %v\n", op, res)
	results <- true
}

var results chan bool = make(chan bool)

func main() {
	fmt.Println("Start")
	n := 10
	funcs := []func(float64, float64) float64{
		fsum,
		fmult,
		math.Max,
	}

	for i := 0; i < n; i++ {
		go send(Op{
			x: rand.Float64() * 10,
			y: rand.Float64() * 10,
			f: funcs[rand.Intn(len(funcs))],
		})
	}
	for i := 0; i < n; i++ {
		go execute()
	}
	for i := 0; i < n; i++ {
		<-results
	}
	fmt.Println("End")
}
