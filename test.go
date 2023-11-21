package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(77))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
}

func add(x int, y int) int {
	return x + y
}
