package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x+y
}

func add_shorter(x,y int) int {
	return x+y
}

func named_return_values(x int) (y, z int) {
	y = x/2
	z = x*2
	return
}

func main() {
	fmt.Printf("Have a number: %g\n", rand.Intn(10))	
	fmt.Printf("Squarepants! %g\n", math.Sqrt(42))
	fmt.Println(math.Pi) 
	fmt.Println(add(23, 19))
	fmt.Println(add_shorter(19, 23))
	fmt.Println(named_return_values(4))
}