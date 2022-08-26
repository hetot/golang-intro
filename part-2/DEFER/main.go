package main

import "fmt"

func main() {
	DefferHelloWorld()
	// DefferHelloWorldWithPanic()
	OrderOfExecution()
	fmt.Println(ChangeWorld())
}

func DefferHelloWorld() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

// Deferred calls are executed even when the function panics
func DefferHelloWorldWithPanic() {
	defer fmt.Println("World")
	panic("Stop")
	fmt.Println("Hello")
}

// Order of execution
func OrderOfExecution() {
	fmt.Println("Hello")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("World")
}

// Use func to return a value
func ChangeWorld() (result string) {
	defer func() {
		result = "Change World" // change value at the very last moment
	}()
	return "Hello World"
}
