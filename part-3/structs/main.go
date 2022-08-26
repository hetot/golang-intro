package main

import (
	"fmt"
	"time"
)

type Person struct {
	Firstname, Lastname string // Compact by combining the various fields of the same type
}

type Content struct {
	Title string
	Body  string
}

type Article struct {
	ID        int
	Content          // Promoted fields
	Author    Person // Nested structs
	CreatedAt *time.Time
}

func main() {
	var a Article // a == Article{ID:0, Content:Content{Title:"", Body:""}, Author:Person{Firstname:"", Lastname:""}, CreatedAt:(*time.Time)(nil)}
	a.Title = "Initialize"
	fmt.Printf("%#v\n", a)

	var b *Article   // b == nil
	b = new(Article) // b == &Article{ID:0, Content:Content{Title:"", Body:""}, Author:Person{Firstname:"", Lastname:""}, CreatedAt:(*time.Time)(nil)}
	b.Title = "Initialize a struct with a struct pointer"
	fmt.Printf("%#v\n", b)

	c := Article{
		Content: Content{
			Title: "Initialize a struct with a struct literal",
		},
	} //  c == Article{ID:0, Content:Content{Title:"Initialize a struct with a struct literal", Body:""}, Author:Person{Firstname:"", Lastname:""}, CreatedAt:(*time.Time)(nil)}
	c.Body = `An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
	An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
	A literal may omit the element list; such a literal evaluates to the zero value for its type.`
	fmt.Printf("%#v\n", c)

	article1 := Article{
		ID: 1,
		Content: Content{
			Title: "Lorem",
			Body:  "Lorem ipsum dolor sit amet. Et molestiae explicabo ut consequatur voluptas nam",
		},
	}
	fmt.Printf("%#v\n", article1.Content)
	fmt.Println("Title:", article1.Title)
	fmt.Println("Body:", article1.Body)

	article2 := struct {
		ID int
		Content
		Author    Person
		CreatedAt *time.Time
	}{ // anonymous structs
		ID: 1,
		Content: Content{
			Title: "Lorem",
			Body:  "Lorem ipsum dolor sit amet. Et molestiae explicabo ut consequatur voluptas nam",
		},
	}

	fmt.Printf("%#v\n", article2.Content)
	fmt.Println("Title:", article2.Title)
	fmt.Println("Body:", article2.Body)

	if article1 == article2 { // Compare structs
		fmt.Println("variables article1 and article2 are equal to each other")
	} else {
		fmt.Println("variables article1 and article2 are NOT equal to each other")
	}

	t1 := time.Now()
	article1.CreatedAt = &t1
	t2 := t1
	article2.CreatedAt = &t2

	if article1 == article2 { // Compare structs
		fmt.Println("variables article1 and article2 are equal to each other")
	} else {
		fmt.Println("variables article1 and article2 are NOT equal to each other")
	}

	fmt.Printf("%#v\n", article1)
	fmt.Printf("%#v\n", article2)

	fmt.Println()

	if *article1.CreatedAt == *article2.CreatedAt {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
