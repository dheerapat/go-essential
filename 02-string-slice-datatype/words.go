package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	fmt.Println(text)

	words := strings.Fields(text)
	fmt.Println(words)

	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	fmt.Println(words)

	count := map[string]int{}

	for _, w := range words {
		_, ok := count[w]
		if !ok {
			count[w] = 1
		} else {
			count[w]++
		}
	}

	fmt.Println(count)

	// alternative solution
	// words := strings.Fields(text)
	// count := map[string]int{}

	// for _, w := range words {
	// 	count[strings.ToLower(w)]++
	// }
}
