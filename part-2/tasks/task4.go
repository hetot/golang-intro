package main

import (
	"fmt"
)

func main() {
	fmt.Println(MySquareRoot(10, 12))
}

func MySquareRoot(num, precision uint) (result float64) {
	// DO NOT USE math.Sqrt!

	result = (float64(num) + float64(num)/float64(num)) / 2

	for result*result > float64(precision) {
		result = (result + float64(num)/result) / 2
	}

	return
}
