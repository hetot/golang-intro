package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var s1 []int = a[1:4]
	fmt.Printf("%v - length of slice %d capacity %d of s1\n", s1, len(s1), cap(s1))

	s2 := []int{6, 7, 8}
	fmt.Printf("%v - length of slice %d capacity %d of s2\n", s2, len(s2), cap(s2))

	s3 := a[:]
	fmt.Printf("%v - length of slice %d capacity %d of s3\n", s3, len(s3), cap(s3))

	s4 := make([]int, 5, 5)
	fmt.Printf("%v - length of slice %d capacity %d of s4\n", s4, len(s4), cap(s4))
}
