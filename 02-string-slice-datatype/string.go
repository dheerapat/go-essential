package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book)) // byte

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// book[0] = 116 string are immutable

	fmt.Println(book[4:11])

	fmt.Println(book[4:])

	fmt.Println(book[:4])

	fmt.Println(book[0:1])
	fmt.Println(book[len(book)-1:])

	fmt.Println("t" + book[1:])

	poem := `
	The sun go down.
	The star come out.
	`

	fmt.Println(poem)
}
