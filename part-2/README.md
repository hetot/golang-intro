
# | Go | Practice 02

[Link to Presentation](https://docs.google.com/presentation/d/11drYotXxtCAxfs5aTEVqg2-o4WVnvYd94Yg22lRaHWo/edit?usp=sharing)

## Topics:
- Packages
- Functions
- FOR
- IF
- SWITCH
- DEFER

## Tasks:
### 1. FizzBuzz
```
package main

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz()
	}
}

func FizzBuzz() {
	//
	// WRITE YOUR CODE HERE
	//
}
```

### 2. Find weekday from given date
```
package main

import (
	"fmt"
	"time"
)

func main() {
	dobStr := "20.04.1998" // Replace this date with your birthday
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %s", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday string) {
	//
	// WRITE YOUR CODE HERE
	//
	return
}
```

### 3. Display numbers  from 1 to 100 in reverse order using DEFER
```
package main

func main() {
	DisplayNumberInReverseOrderWithDefer()
}

func DisplayNumberInReverseOrderWithDefer() {
	for i := 0; i < 100; i++ {
		//
		// WRITE YOUR CODE HERE
		//
	}
}
```

### 4. Write a function to calculate square root of a given number upto a given precision
```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(MySquareRoot(10, 12))
}

func MySquareRoot(num, precision uint) (result float64) {
	// DO NOT USE math.Sqrt!

	//
	// WRITE YOUR CODE HERE
	//

	return
}
```

### 5. Find the Minimum Number
Link: https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number/problem
```
package main

func main() {
	n := 12
	// Read n from input
	DisplayMinimumNumberFunction(n)
}

// https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number
func DisplayMinimumNumberFunction(n int) {
	//
	// WRITE YOUR CODE HERE
	//
}
```

## Reference
- https://blog.golang.org/using-go-modules
- https://www.callicoder.com/golang-packages
- https://zetcode.com/golang/function
- https://yourbasic.org/golang/for-loop
- https://yourbasic.org/golang/if-else-statement
- https://golangbot.com/if-statement
- https://www.geeksforgeeks.org/switch-statement-in-go
- https://yourbasic.org/golang/switch-statement
- https://yourbasic.org/golang/defer
- https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number/problem

## FAQ

#### Question 1

Answer 1

#### Question 2

Answer 2

  
## Feedback
If you have any feedback, please reach out to me at saidamir.botirov@gmail.com
## Authors
[@Saidamir_Botirov](https://www.linkedin.com/in/saidamir-botirov)
