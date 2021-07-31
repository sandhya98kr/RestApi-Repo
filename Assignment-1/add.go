package main

import "fmt"

//creating a variable of type string
var one string

//function to reverse the slice num2[]
func rev2(x []int) []int {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	return x
}
func main() {
	//creating slices
	var num1 []int = []int{1, 2, 3, 4, 5, 6}
	var num2 []int = []int{1, 2, 3, 4, 5, 6}
	//calling a function to reverse a slice
	rev2(num2)
	//Appending slices
	z := append([]int{}, append(num1, num2...)...)
	//calling add function
	add(z)

}

//creating a function add and it returns string
func add(z []int) string {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Printf("Error is :%v\n", r)
		}
	}()
	for i := 0; i < len(z); i++ {
		//logic to add
		one := z[i] + z[i+6]
		if i == 6 {
			break
		}
		fmt.Print("\t", one)
	}
	return one
}
