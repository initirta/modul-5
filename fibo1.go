package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a)
	b = 0
	c = 1

	if a >= 0 {
		fmt.Print(b, " ")
	}
	if a >= 1 {
		fmt.Print(c, " ")
	}

	for i := 0; i < a; i++ {
		d = b
		b = c
		c = c + d
		fmt.Print(c, " ")

	}

}
