package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// Work item struct: two operands and a function object
type Op struct {
	x  float64
	y  float64
	fn Function
}

// Function object: name and func
type Function struct {
	name string
	impl func(float64, float64) float64
}

// Pipeline channel
var c chan Op = make(chan Op)

var results chan bool = make(chan bool)

// WaitGroup to ensure main
var wg sync.WaitGroup

func main() {
	fmt.Println("Start")
	n := 10
	// Available functions, some from standard library, some custom
	funcs := []Function{
		Function{"min", math.Min},
		Function{"max", math.Max},
		Function{"foo", func(x, y float64) float64 {
			return x*x + y*y
		}},
	}

	// Send random Ops to channel
	for i := 0; i < n; i++ {
		// Inline goroutine to write Op to channel without blocking
		// Op function and operands are random
		go func(op Op) {
			c <- op
		}(Op{
			x:  rand.Float64() * 10,
			y:  rand.Float64() * 10,
			fn: funcs[rand.Intn(len(funcs))],
		})
		//time.Sleep(time.Duration(1) * time.Second)
	}
	// Execute all Ops from the channel
	for i := 0; i < n; i++ {
		op := <-c
		res := op.fn.impl(op.x, op.y)
		fmt.Printf("%s(%0.2f,%0.2f)= %0.2f\n", op.fn.name, op.x, op.y, res)
	}
	fmt.Println("End")
}
