// Go program to illustrate the
// use of assignment operators
package main

import "fmt"

func main() {
	var p int = 45
	var q int = 50
	fmt.Printf("p = %d,\tq = %d \n", p, q)
	// “=”(Simple Assignment)
	p = q
	fmt.Println("p  = q,\tp =", p)
	// “+=”(Add Assignment)
	p += q
	fmt.Println("p += q,\tp =", p)
	//“-=”(Subtract Assignment)
	p -= q
	fmt.Println("p -= q,\tp =", p)
	// “*=”(Multiply Assignment)
	p *= q
	fmt.Println("p *= q,\tp =", p)
	// “/=”(Division Assignment)
	p /= q
	fmt.Println("p /= q,\tp =", p)
	// “%=”(Modulus Assignment)
	p %= q
	fmt.Println("p %= q,\tp =", p)
}
