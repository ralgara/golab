package main

import (
	"fmt"
	"testing"
)

func TestAbsVector(t *testing.T) {
	v := Vector{3, 4}
	got := v.Abs()
	want := 5.
	if got != want {
		t.Errorf("Abs(%v) = %.2f; want %.2f", v, got, want)
	}
}

func ExampleFoobar() {
	fmt.Println(len("foobar"))
	// Output: 6
}
