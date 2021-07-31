package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go EvenNumbers()
	go OddNumbers()
	fmt.Println("invoked main routine")
	wg.Wait()

}
func EvenNumbers() {
	defer wg.Done()
	fmt.Println("\nEven Numbers from 1 to 30.........")
	num := 30
	for i := 2; i <= num; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\t", i)
		}
	}
}
func OddNumbers() {
	defer wg.Done()
	fmt.Println("\noddnumbers from 1 to 30.....")
	num := 30
	for i := 2; i <= num; i++ {
		if i%2 != 0 {
			fmt.Printf("%d\t", i)
		}
	}
}
