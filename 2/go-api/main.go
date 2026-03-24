package main

import "fmt"

func main() {
	x := 20
	updateVal(&x)
	fmt.Print(x)
}
func updateVal(x *int) {
	*x = *x + 50
}
