package samples

import (
	"math/rand"
	"time"
)

// Go sample function
func Add(a int, b int) int {
	return a + b
}

// Go function omit type
func Sub(x, y int) int {
	return x - y
}

// Go function named return variables
func Inc(x, y, z int) (a, b, c int) {
	a = x + 1
	b = y + 1
	c = z + 1
	return
}

// Go function multiple return values
func Threerandom() (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	z := rand.Intn(10)
	return x, y, z
}

// Go variadic function
func Sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// Go recursive function
func Fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Fact(n-1)
}
