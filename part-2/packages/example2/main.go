package main

import (
	"fmt"
	str "strings" // Package Alias

	"example2/numbers"
	"example2/strings"
	"example2/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)
	fmt.Println(greeting.MorningText)
	fmt.Println(greeting.EveningText)
	// fmt.Println(greeting.loremIpsumText) // cannot refer to unexported name greeting.loremIpsumTex

	fmt.Println(strings.Reverse("callicoder"))
	// fmt.Println(string(strings.reverseRunes([]rune("callicoder")))) // cannot refer to unexported name strings.reverseRunes

	fmt.Println(str.Count("Go is Awesome. Go is Fun. I love Coding", "Go"))
}
