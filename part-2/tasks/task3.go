package main

import (
	"fmt"
)

func main() {
	DisplayNumberInReverseOrderWithDefer()
}

func DisplayNumberInReverseOrderWithDefer() {
	for i := 0; i < 100; i++ {
		defer displayNum(i)
	}
}

func displayNum(x int) {
	fmt.Println(x)
}
