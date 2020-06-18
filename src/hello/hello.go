package main

import (
	"fmt"
	"strings"
	"math"
)

type Shape struct {
	height int
	width int
}

func hello(i int) int {
	return i*2
}

func main() {
	types_and_pointers()
	structs()
	arrays()
	maps()
	funcs()
	interfaces()
}

type MyFloat float64
type Abser interface {
	Abs() float64
}

func (f MyFloat) Abs() float64 {
	return math.Abs(float64(f))
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type Vertex struct {
	X, Y float64
}

func interfaces() {
	fmt.Println("\n=== Interfaces")
	fmt.Println("= Methods")
	var a Abser = MyFloat(-3.1416)
	fmt.Printf("Abs(MyFloat(%v)): %v\n", a, a.Abs())
	var v Abser = Vertex{X:1, Y:1}
	fmt.Printf("Abs(Vertex{%v}): %v\n", v, v.Abs())

	fmt.Println("= The empty interface")
	var describe func(interface{}) string = func(i interface{}) string {
		return fmt.Sprintf("(%v, %T)", i, i)
	}
	var i interface{}
	fmt.Println("describe empty interface set to nil: ", describe(i))
	i = 32
	fmt.Println("describe empty interface set to int(32): ", describe(i))
	i = "foo"
	fmt.Println("describe empty interface set to 'foo': ", describe(i))
}

func types_and_pointers() {
	x, y := 7,11
	z := &x
	var p int = *z
	fmt.Print(x, y, *z, hello(p), "\n")
}

func structs() {
	fmt.Printf("\n===Structs\nInitialization\n")
	s := Shape{1,2}
	fmt.Printf("Field access\n")
	s.height = hello(s.width)
	fmt.Printf("Init with named fields and defaults")
	t := Shape{width: 7}
	fmt.Print(s, t, "\n")
}

func arrays() {
	fmt.Printf("\n=== Arrays and slices")
	primes := []int{1,2,3,5,7,11,13,17}
	primes[3] = 23
	var otherPrimes = &primes
	slice := primes[3:4]
	fmt.Print(primes, (*otherPrimes)[1:4], slice, "\n")
	fmt.Printf("Length and capacity\n")
	fmt.Print(len(slice), cap(slice), "\n")
	fmt.Printf("Creation\n")
	m := make([]int, 10)
	fmt.Printf("%v\n", m)
	fmt.Printf("Slices of slices\n")
	board := [][]string{
		[]string{"x", "_", "_"},
		[]string{"_", "o", "_"},
		[]string{"_", "_", "x"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("Append and range")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	pow = append(pow, 256,512,1024)
	for i, v := range pow {
		fmt.Println(i, v)
	}
}

func maps() {
	fmt.Println("\n=== Maps")
	fmt.Println("Literals")
	m := make(map[string]Shape)
	m["square"] = Shape{height: 1, width: 1}
	m["rectangle"] = Shape{3,1}
	n := map[string]Shape{
		"foo": Shape{
			17,
			41,
		},
		"bar": Shape{
			height: 1,
		},
	}
	fmt.Println(m, n)
	fmt.Println("Deletion")
	delete(n, "bar")
	fmt.Println(n)
	fmt.Println("Access with check")
	v, ok := n["foo"]
	fmt.Println(v, ok)
	v, ok = n["bar"]
	fmt.Println(v, ok)
}

func funcs() {
	fmt.Println("\n=== Functions")
	f := func (fn func (string, string) string) string {
		return fn("foo", "bar")
	}

	g := func (s string, t string) string {
		return fmt.Sprintf("%s is a %s", s, t)
	}

	fmt.Println(f(g))
}