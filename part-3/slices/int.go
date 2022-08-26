package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var s1 []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(s1)

	s2 := []int{6, 7, 8} // creates array and returns a slice reference
	fmt.Println(s2)

	s3 := a[:] // creates a slice which contains all elements of the a
	fmt.Println(s3)

	s4 := make([]int, 5, 5) // The values are zeroed by default when a slice is created using make
	fmt.Println(s4)
}
