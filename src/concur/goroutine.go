package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fictional work item: an ID with a `magic` number resulting from titanic effort (random wait)
type Item struct {
	id    int
	magic int
}

// Worker goroutine: get an ID, does some pretend work (random) and writes result to channel
func worker(id int) int {
	delay := rand.Float64() * 3
	time.Sleep(time.Duration(delay) * time.Second)
	magic := rand.Intn(1000)
	item := Item{id: id, magic: magic}
	channel <- item
	fmt.Printf("Sent [%v] after %0.2f seconds\n", item, delay)
	return magic
}

// Channel connecting workers and main
var channel chan Item = make(chan Item, 10)

func main() {
	N := 10
	for i := 0; i < N; i++ {
		go worker(i)
	}
	for i := 0; i < N; i++ {
		m := <-channel
		fmt.Printf("Received %v\n", m)
	}
	fmt.Println("Done")
}
