package main

import "fmt"

func main() {
	var x int = 100
	var z int = x
	fmt.Printf("&x=%p, x=%d\n", &x, x)
	fmt.Printf("&z=%p, z=%d\n", &z, z)

	var y *int = &x
	fmt.Printf("y=%p, *y=%d\n", y, *y)
	fmt.Println("&y = ", &y)

	*y++
	fmt.Printf("&x=%p, x=%d\n", &x, x)
	fmt.Printf("&z=%p, z=%d\n", &z, z)

	y = &z
	*y = -100
	fmt.Printf("&x=%p, x=%d\n", &x, x)
	fmt.Printf("&z=%p, z=%d\n", &z, z)
}
