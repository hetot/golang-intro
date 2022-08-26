package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}

	fmt.Println("Simple loop:")
	for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

	fmt.Println("Range loop:")
	for i, v := range a { // range returns both the index and value
		fmt.Printf("%d th element of a is %.2f\n", i, v)
	}

	fmt.Println("Range loop while ommiting index:")
	var i uint
	for _, v := range a { // range returns both the index and value
		i++
		fmt.Printf("%d th element of a is %.2f\n", i, v)
	}

	fmt.Println("Range loop while ommiting value:")
	for i, _ := range a { // range returns both the index and value
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}
