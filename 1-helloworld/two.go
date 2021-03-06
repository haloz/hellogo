package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	var z int
	z = x+y
	return z	
}

func add_shorter(x,y int) int {
	return x+y
}

func named_return_values(x int) (y, z int) {
	var d, e int = 2,2
	y = x/d
	z = x*e
	return
}

func basictypes() {
	var (
		t bool = true
		MaxInt uint64 = 1<<64-1
	)
	fmt.Printf("type: %T, value: %v\n", t,t)
	fmt.Printf("type: %T, value: %v\n", MaxInt,MaxInt)
}

func main() {
	s := "number" // short variable declaration
	fmt.Printf("Have a %s: %g\n", s, rand.Intn(10))	
	fmt.Printf("Squarepants! %g\n", math.Sqrt(42))
	fmt.Println(math.Pi) 		
	fmt.Println(add(23, 19))
	fmt.Println(add_shorter(19, 23))
	fmt.Println(named_return_values(4))	
	basictypes()

	var f float64 = math.Sqrt(55)
	fmt.Println(f)
	var fr = uint(f)
	fmt.Println(fr)
}