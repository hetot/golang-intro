package main

import (
	"fmt"
	"lesson-02/functions/samples"
)

func main() {
	x := 4
	y := 5

	z := samples.Add(x, y)
	f := samples.Add

	fmt.Println(z)       // 9
	fmt.Println(f(x, y)) // 9

	fmt.Println(samples.Sub(x, y)) // -1

	fmt.Println(samples.Inc(x, y, z)) // 5 6 10

	fmt.Println(samples.Threerandom()) // 3 random numbers

	// Go anonymous function
	sum := func(a, b, c int) int {
		return a + b + c
	}(3, 5, 7)
	fmt.Println(sum) // 15

	fmt.Println(samples.Sum(1, 2, 3))       // 6
	fmt.Println(samples.Sum(1, 2, 3, 4))    // 10
	fmt.Println(samples.Sum(1, 2, 3, 4, 5)) // 15

	fmt.Println(samples.Fact(7))  // 5040
	fmt.Println(samples.Fact(10)) // 3628800
	fmt.Println(samples.Fact(15)) // 1307674368000
}
