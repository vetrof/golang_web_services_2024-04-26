package main

import "fmt"

func main() {
	x := 10
	n := 0
	for range x {

		fmt.Println(n)
		n += 1
	}
}
