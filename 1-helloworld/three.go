package main

import (
	"fmt"
	"math"
)

func loops() (output string) {
	// normal for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum+=i
		output += fmt.Sprintf("%v\n", sum)	
	}	

	// optional init and post statement
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}	
	output += fmt.Sprintf("%v\n", sum)

	// no while, just use for - perl-like
	sum = 0
	for sum < 10 {
		sum++
	}
	output += fmt.Sprintf("%v\n", sum)
	return
}

func ifs() (output string) {
	sum := 0
	for sum < 10 {
		sum++
		if sum % 2 == 1 {
			output += "blubb\n"
		} else {
			output += "tick\n"
		}
	}

	if v := math.Pow(2,2); v < 4 {
		output += fmt.Sprintf("%v\n", v)
	} else {
		v = v-1; // can use v here
		output += fmt.Sprintf("%v\n", v)
	}
	// can't use v here, only valid in scope of the if-else statement 
	return
}

func main() {
	fmt.Println(loops())
	fmt.Println(ifs())
}
