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

	a, b := 0.0, 1.0 // Пределы интегрирования
	n := 10000       // Количество прямоугольников
	result := integrate(a, b, n)
	fmt.Printf("Результат интегрирования: %f\n", result)
}

func add(x int, y int) int {
	return x + y
}

func f(x float64) float64 {
	return x * x // Функция, которую нужно проинтегрировать
}

func integrate(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return sum * h
}
