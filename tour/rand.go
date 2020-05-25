package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("My favorite number is %d \n", rand.Intn(10))
	fmt.Printf("Square root of 9 is %g\nSquare root of 7 is %g\n",
		math.Sqrt(9), math.Sqrt(7))
}
