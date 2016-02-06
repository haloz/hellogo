package main

import (
	"fmt"
)

func newton_sqrt(x float64) float64 {
	z := 1.0
	
	for i := 0; i < 100; i++ {
		z = z - (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("sqrt of %v is %v\n",i,newton_sqrt(float64(i)))
	}	

}