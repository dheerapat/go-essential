package main

import (
	"fmt"
)

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %q (type of %T)\n", s, s)

	count := 0

	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			r := a * b
			sr := fmt.Sprintf("%d", r)
			if sr[0] == sr[len(sr)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
