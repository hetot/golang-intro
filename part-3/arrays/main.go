package main

import (
	"fmt"
)

func main() {
	var a1 [3]int   // int array with length 3
	fmt.Println(a1) // All elements in an array are automatically assigned the zero value of the array type
	a1[0] = 12      // array index starts at 0
	a1[1] = 78
	a1[2] = 50
	fmt.Println(a1)

	a2 := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a2)

	a3 := [3]int{12} // It is not necessary that all elements in an array have to be assigned
	fmt.Println(a3)

	a4 := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a4)

	/*
		b1 := [3]int{5, 78, 8}
		var b2 [5]int
		b1 = b2 // not possible since [3]int and [5]int are distinct types
	*/

	c1 := [...]string{"USA", "China", "India", "Germany", "France"}
	c2 := c1 // a copy of c1 is assigned to c2
	c1[0] = "Singapore"
	fmt.Println("c1 is ", c1)
	fmt.Println("c2 is ", c2)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) // num is passed by value
	fmt.Println("after passing to function ", num)

	fmt.Println("length of num is", len(num))
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
