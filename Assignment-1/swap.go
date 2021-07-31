package main

import "fmt"

func main() {
	x := 2
	y := 3
	fmt.Printf("before swaping value of x is %d and y is %d", x, y)
	swap(&x, &y)
	fmt.Printf("after swapping value of x is %d and y is %d", x, y)

}
func swap(a, b *int) {
	tempx := *a
	tempy := *b
	*a = tempy
	*b = tempx
}

//used go build -gcflags="-m" command for Escape analysis
