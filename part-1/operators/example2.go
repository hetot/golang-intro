// Go program to illustrate the
// use of relational operators
package main

import "fmt"

func main() {
	p := 34
	q := 20
	fmt.Printf("p = %d, q = %d \n", p, q)
	// ‘=='(Equal To)
	fmt.Println("Result of (p == q) = ", p == q)

	// ‘!='(Not Equal To)
	fmt.Println("Result of (p != q) = ", p != q)

	// ‘<‘(Less Than)
	fmt.Println("Result of (p < q) = ", p < q)

	// ‘>'(Greater Than)
	fmt.Println("Result of (p > q) = ", p > q)

	// ‘>='(Greater Than Equal To)
	fmt.Println("Result of (p >= q) = ", p >= q)

	// ‘<='(Less Than Equal To)
	fmt.Println("Result of (p <= q) = ", p <= q)
}
