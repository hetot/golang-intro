package main

import (
	"fmt"
)

func main() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("Array before change 1", numa)
	nums1[0] = 100
	fmt.Println("Array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	fmt.Println()

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("Array before:", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("Array after:", darr)

	fmt.Println()

	nos := []int{8, 7, 6}
	fmt.Println("Slice before function call", nos)
	subtactOne(nos)                               // function modifies the slice
	fmt.Println("Slice after function call", nos) // modifications are visible outside
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 1
	}
}
